/* global stripe */
import React, {
  useState,
  useCallback,
  useMemo,
  useContext,
  useEffect,
} from "react";
import cn from "classnames";
import { useParams } from "react-router-dom";
import { DateTime } from "luxon";
import {
  Text,
  DefaultEffects,
  Dialog,
  DialogType,
  DialogFooter,
  IDialogContentProps,
  ThemeProvider,
  PartialTheme,
  Spinner,
  SpinnerSize,
  Pivot,
  PivotItem,
} from "@fluentui/react";
import { useConst } from "@fluentui/react-hooks";
import { Context, FormattedMessage } from "@oursky/react-messageformat";
import ScreenTitle from "../../ScreenTitle";
import ShowError from "../../ShowError";
import ShowLoading from "../../ShowLoading";
import {
  Subscription,
  SubscriptionPlan,
  SubscriptionUsage,
} from "./globalTypes.generated";
import { PortalAPIAppConfig } from "../../types";
import { AppFragmentFragment } from "./query/subscriptionScreenQuery.generated";
import { useSubscriptionScreenQueryQuery } from "./query/subscriptionScreenQuery";
import styles from "./SubscriptionScreen.module.css";
import { useLoading, useIsLoading } from "./../../hook/loading";
import ButtonWithLoading from "../../ButtonWithLoading";
import { useSetSubscriptionCancelledStatusMutation } from "./mutations/setSubscriptionCancelledStatusMutation";
import { useSystemConfig } from "../../context/SystemConfigContext";
import ErrorDialog from "../../error/ErrorDialog";
import ScreenLayoutScrollView from "../../ScreenLayoutScrollView";
import PrimaryButton from "../../PrimaryButton";
import DefaultButton from "../../DefaultButton";
import { useCancelFailedSubscriptionMutation } from "./mutations/cancelFailedSubscriptionMutation";
import ExternalLink from "../../ExternalLink";
import { isStripePlan, Plan, CTAVariant } from "../../util/plan";
import {
  PlanCardBusiness,
  PlanCardDevelopers,
  PlanCardEnterprise,
  PlanCardFree,
} from "../../components/billing/PlanCard";
import { useCreateCheckoutSessionMutation } from "./mutations/createCheckoutSessionMutation";
import { useUpdateSubscriptionMutation } from "./mutations/updateSubscriptionMutation";
import { usePreviewUpdateSubscriptionMutation } from "./mutations/previewUpdateSubscriptionMutation";
import { formatDateOnly } from "../../util/formatDateOnly";
import { FeatureBanner } from "../../components/billing/FeatureBanner";
import ScreenDescription from "../../ScreenDescription";

const CHECK_IS_PROCESSING_SUBSCRIPTION_INTERVAL = 5000;

const CONTACT_US_BUTTON_THEME: PartialTheme = {
  palette: {
    themePrimary: "#c8c8c8",
    neutralPrimary: "#c8c8c8",
  },
  semanticColors: {
    linkHovered: "#c8c8c8",
  },
};

function PlansSection({
  currentPlanName,
  subscriptionCancelled,
  nextBillingDate,
  subscriptionPlans,
  onClickContactUs,
  onClickCancelSubscription,
}: {
  currentPlanName: string;
  subscriptionCancelled: boolean;
  nextBillingDate: Date | undefined;
  subscriptionPlans: SubscriptionPlan[];
  onClickContactUs: () => void;
  onClickCancelSubscription: () => void;
}) {
  const {
    themes: { destructive },
  } = useSystemConfig();
  const { locale } = useContext(Context);
  const [upgradeToPlan, setUpgradeToPlan] = useState<string | null>(null);
  const [downgradeToPlan, setDowngradeToPlan] = useState<string | null>(null);
  const [isReactiveDialogHidden, setIsReactiveDialogHidden] =
    useState<boolean>(true);
  const { appID } = useParams() as { appID: string };
  const { createCheckoutSession, loading: createCheckoutSessionLoading } =
    useCreateCheckoutSessionMutation();
  useLoading(createCheckoutSessionLoading);
  const [updateSubscription, { loading: updateSubscriptionLoading }] =
    useUpdateSubscriptionMutation();
  useLoading(updateSubscriptionLoading);
  const {
    setSubscriptionCancelledStatus,
    loading: reactivateSubscriptionLoading,
    error: reactivateSubscriptionError,
  } = useSetSubscriptionCancelledStatusMutation(appID);
  useLoading(reactivateSubscriptionLoading);

  const [previewUpdateSubscription, { data, loading }] =
    usePreviewUpdateSubscriptionMutation();
  useLoading(loading);

  const onClickSubscribe = useCallback(
    (planName: string) => {
      createCheckoutSession(appID, planName)
        .then((url) => {
          if (url) {
            window.location.href = url;
          }
        })
        .finally(() => {});
    },
    [appID, createCheckoutSession]
  );

  const onConfirmUpgrade = useCallback(() => {
    if (!upgradeToPlan) {
      console.error("upgradeToPlan should not be null");
      return;
    }
    updateSubscription({
      variables: {
        appID,
        planName: upgradeToPlan,
      },
    }).finally(() => {
      setUpgradeToPlan(null);
    });
  }, [appID, updateSubscription, upgradeToPlan]);

  const onConfirmDowngrade = useCallback(() => {
    if (!downgradeToPlan) {
      console.error("downgradeToPlan should not be null");
      return;
    }
    updateSubscription({
      variables: {
        appID,
        planName: downgradeToPlan,
      },
    }).finally(() => {
      setDowngradeToPlan(null);
    });
  }, [appID, downgradeToPlan, updateSubscription]);

  const onClickUpgrade = useCallback(
    (planName: string) => {
      previewUpdateSubscription({
        variables: {
          appID,
          planName,
        },
      }).finally(() => {});
      setUpgradeToPlan(planName);
    },
    [appID, previewUpdateSubscription]
  );

  const onClickDowngrade = useCallback(
    (planName: string) => {
      previewUpdateSubscription({
        variables: {
          appID,
          planName,
        },
      }).finally(() => {});
      setDowngradeToPlan(planName);
    },
    [appID, previewUpdateSubscription]
  );

  const onClickReactivate = useCallback(() => {
    setIsReactiveDialogHidden(false);
  }, []);

  const onClickConfirmReactivate = useCallback(async () => {
    try {
      await setSubscriptionCancelledStatus(false);
    } finally {
      setIsReactiveDialogHidden(true);
    }
  }, [setSubscriptionCancelledStatus]);

  const onPlanAction = useMemo(() => {
    const plans: Plan[] = ["enterprise"];
    if (subscriptionPlans.findIndex((p) => p.name === "developers") !== -1) {
      plans.push("developers");
    }
    if (subscriptionPlans.findIndex((p) => p.name === "business") !== -1) {
      plans.push("business");
    }
    plans.push("enterprise");

    return plans.reduce<Partial<Record<Plan, (action: CTAVariant) => void>>>(
      (callbacks, plan) => {
        const fn = (action: CTAVariant) => {
          switch (action) {
            case "subscribe":
              onClickSubscribe(plan);
              break;
            case "upgrade":
              onClickUpgrade(plan);
              break;
            case "downgrade":
              onClickDowngrade(plan);
              break;
            case "reactivate":
              onClickReactivate();
              break;
            case "contact-us":
              onClickContactUs();
              break;
            case "current":
            case "non-applicable":
            default:
              console.error(
                `action button clicked but action:${action} should not be clickable. plan: ${plan}`
              );
              break;
          }
        };
        callbacks[plan] = fn;
        return callbacks;
      },
      {}
    );
  }, [
    onClickContactUs,
    onClickDowngrade,
    onClickReactivate,
    onClickSubscribe,
    onClickUpgrade,
    subscriptionPlans,
  ]);

  const onFreePlanAction = useCallback(
    (action: CTAVariant) => {
      switch (action) {
        case "downgrade":
          // Downgrade to free plan means cancel subcription
          onClickCancelSubscription();
          break;
        // All other cases should not happen
        default:
          console.error(
            `action button clicked but action:${action} should not be clickable. plan: free`
          );
          break;
      }
    },
    [onClickCancelSubscription]
  );

  const amountDue =
    data?.previewUpdateSubscription.amountDue != null
      ? data.previewUpdateSubscription.amountDue / 100
      : null;
  const formattedDate = formatDateOnly(locale, nextBillingDate ?? null);

  // @ts-expect-error
  const upgradeDialogContentProps: IDialogContentProps = useMemo(() => {
    return {
      type: DialogType.normal,
      title: <FormattedMessage id="SubscriptionScreen.upgrade.title" />,
      subText:
        amountDue == null ? (
          <FormattedMessage id="loading" />
        ) : (
          <FormattedMessage
            id="SubscriptionScreen.upgrade.description"
            values={{
              amount: amountDue,
              date: formattedDate ?? "",
            }}
          />
        ),
    };
  }, [amountDue, formattedDate]);

  // @ts-expect-error
  const downgradeDialogContentProps: IDialogContentProps = useMemo(() => {
    return {
      type: DialogType.normal,
      title: <FormattedMessage id="SubscriptionScreen.downgrade.title" />,
      subText:
        amountDue == null ? (
          <FormattedMessage id="loading" />
        ) : (
          <FormattedMessage
            id="SubscriptionScreen.downgrade.description"
            values={{
              amount: amountDue,
              date: formattedDate ?? "",
            }}
          />
        ),
    };
  }, [amountDue, formattedDate]);

  // @ts-expect-error
  const reactivateDialogContentProps: IDialogContentProps = useMemo(() => {
    return {
      type: DialogType.normal,
      title: <FormattedMessage id="SubscriptionScreen.reactivate.title" />,
      subText: (
        <FormattedMessage id="SubscriptionScreen.reactivate.confirmation" />
      ),
    };
  }, []);

  const isLoading = useIsLoading();

  const onDismissUpgradeDialog = useCallback(() => {
    setUpgradeToPlan(null);
  }, []);

  const onDismissDowngradeDialog = useCallback(() => {
    setDowngradeToPlan(null);
  }, []);

  const onDismissReactiveDialog = useCallback(() => {
    setIsReactiveDialogHidden(true);
  }, []);

  return (
    <>
      <div className="overflow-x-auto w-full">
        <div className="grid grid-flow-col grid-rows-1 auto-cols-[1fr] gap-4">
          <PlanCardFree
            currentPlan={currentPlanName}
            subscriptionCancelled={subscriptionCancelled}
            onAction={onFreePlanAction}
          />
          {onPlanAction.developers != null ? (
            <PlanCardDevelopers
              currentPlan={currentPlanName}
              subscriptionCancelled={subscriptionCancelled}
              onAction={onPlanAction.developers}
            />
          ) : null}
          {onPlanAction.business != null ? (
            <PlanCardBusiness
              currentPlan={currentPlanName}
              subscriptionCancelled={subscriptionCancelled}
              onAction={onPlanAction.business}
            />
          ) : null}
          <PlanCardEnterprise
            currentPlan={currentPlanName}
            subscriptionCancelled={subscriptionCancelled}
            onAction={onPlanAction.enterprise!}
          />
        </div>
      </div>
      <ErrorDialog
        error={reactivateSubscriptionError}
        rules={[]}
        fallbackErrorMessageID="SubscriptionScreen.reactivate.error"
      />
      <Dialog
        hidden={upgradeToPlan == null}
        onDismiss={onDismissUpgradeDialog}
        dialogContentProps={upgradeDialogContentProps}
      >
        <DialogFooter>
          <PrimaryButton
            onClick={onConfirmUpgrade}
            disabled={isLoading}
            text={<FormattedMessage id="SubscriptionScreen.label.upgrade" />}
          />
          <DefaultButton
            onClick={onDismissUpgradeDialog}
            text={<FormattedMessage id="cancel" />}
          />
        </DialogFooter>
      </Dialog>
      <Dialog
        hidden={downgradeToPlan == null}
        onDismiss={onDismissDowngradeDialog}
        dialogContentProps={downgradeDialogContentProps}
      >
        <DialogFooter>
          <PrimaryButton
            onClick={onConfirmDowngrade}
            theme={destructive}
            disabled={isLoading}
            text={<FormattedMessage id="SubscriptionScreen.label.downgrade" />}
          />
          <DefaultButton
            onClick={onDismissDowngradeDialog}
            text={<FormattedMessage id="cancel" />}
          />
        </DialogFooter>
      </Dialog>
      <Dialog
        hidden={isReactiveDialogHidden}
        onDismiss={onDismissReactiveDialog}
        dialogContentProps={reactivateDialogContentProps}
      >
        <DialogFooter>
          <ButtonWithLoading
            loading={reactivateSubscriptionLoading}
            onClick={onClickConfirmReactivate}
            disabled={isReactiveDialogHidden}
            labelId="confirm"
          />
          <DefaultButton
            onClick={onDismissReactiveDialog}
            disabled={isReactiveDialogHidden || reactivateSubscriptionLoading}
            text={<FormattedMessage id="cancel" />}
          />
        </DialogFooter>
      </Dialog>
    </>
  );
}

interface SubscriptionScreenContentProps {
  appID: string;
  planName: string;
  subscription?: Subscription;
  subscriptionPlans: SubscriptionPlan[];
  thisMonthUsage?: SubscriptionUsage;
  previousMonthUsage?: SubscriptionUsage;
  effectiveAppConfig?: PortalAPIAppConfig;
}

enum Tab {
  Subscription = "Subscription",
  PlanDetail = "PlanDetail",
}

function SubscriptionScreenContent(props: SubscriptionScreenContentProps) {
  const { appID, planName, subscription, subscriptionPlans, thisMonthUsage } =
    props;
  const { themes } = useSystemConfig();
  const { renderToString } = useContext(Context);

  const subscriptionCancelled = useMemo(() => {
    return !!subscription?.endedAt;
  }, [subscription?.endedAt]);

  const nextBillingDate = useMemo(() => {
    if (!isStripePlan(planName)) {
      return undefined;
    }

    const nextBillingDate = thisMonthUsage?.nextBillingDate;
    if (nextBillingDate != null) {
      return new Date(nextBillingDate);
    }
    return undefined;
  }, [planName, thisMonthUsage]);

  const [enterpriseDialogHidden, setEnterpriseDialogHidden] = useState(true);
  const [cancelDialogHidden, setCancelDialogHidden] = useState(true);

  const [selectedTab, setSelectedTab] = useState<Tab>(Tab.Subscription);
  const onTabChange = useCallback((item?: PivotItem) => {
    if (item == null) {
      return;
    }
    const { itemKey } = item.props;
    setSelectedTab(itemKey as Tab);
  }, []);

  const enterpriseDialogContentProps: IDialogContentProps = useMemo(() => {
    return {
      type: DialogType.normal,
      title: <FormattedMessage id="SubscriptionScreen.enterprise.title" />,
      // @ts-expect-error
      subText: (
        <FormattedMessage id="SubscriptionScreen.enterprise.instructions" />
      ) as IDialogContentProps["subText"],
    };
  }, []);

  const cancelDialogContentProps: IDialogContentProps = useMemo(() => {
    if (!subscription) {
      return {
        type: DialogType.normal,
        title: <FormattedMessage id="SubscriptionScreen.cancel.title" />,
        // @ts-expect-error
        subText: (
          <FormattedMessage id="SubscriptionScreen.cancel.confirmation.customPlan" />
        ) as IDialogContentProps["subText"],
      };
    }
    return {
      type: DialogType.normal,
      title: <FormattedMessage id="SubscriptionScreen.cancel.title" />,
      // @ts-expect-error
      subText: (
        <FormattedMessage id="SubscriptionScreen.cancel.confirmation" />
      ) as IDialogContentProps["subText"],
    };
  }, [subscription]);

  const onClickContactUs = useCallback(() => {
    setEnterpriseDialogHidden(false);
  }, []);

  const onClickCancel = useCallback((e?: React.MouseEvent) => {
    e?.preventDefault();
    e?.stopPropagation();
    setCancelDialogHidden(false);
  }, []);

  const onDismiss = useCallback(() => {
    setEnterpriseDialogHidden(true);
    setCancelDialogHidden(true);
  }, []);

  const {
    setSubscriptionCancelledStatus,
    loading: cancelSubscriptionLoading,
    error: cancelSubscriptionError,
  } = useSetSubscriptionCancelledStatusMutation(appID);
  useLoading(cancelSubscriptionLoading);

  const onClickCancelSubscriptionConfirm = useCallback(
    (e) => {
      e.preventDefault();
      e.stopPropagation();
      setSubscriptionCancelledStatus(true)
        .catch(() => {})
        .finally(() => {
          onDismiss();
        });
      setCancelDialogHidden(false);
    },
    [setSubscriptionCancelledStatus, onDismiss, setCancelDialogHidden]
  );

  return (
    <>
      <Dialog
        hidden={cancelDialogHidden}
        onDismiss={onDismiss}
        dialogContentProps={cancelDialogContentProps}
      >
        <DialogFooter>
          {!!subscription ? (
            <ButtonWithLoading
              theme={themes.destructive}
              loading={cancelSubscriptionLoading}
              onClick={onClickCancelSubscriptionConfirm}
              disabled={cancelDialogHidden}
              labelId="confirm"
            />
          ) : (
            <PrimaryButton
              href="mailto:hello@authgear.com"
              onClick={onDismiss}
              disabled={cancelDialogHidden}
              text={
                <FormattedMessage id="SubscriptionScreen.cancel.confirmation.customPlan.button" />
              }
            />
          )}
          <DefaultButton
            onClick={onDismiss}
            disabled={cancelSubscriptionLoading || cancelDialogHidden}
            text={<FormattedMessage id="cancel" />}
          />
        </DialogFooter>
      </Dialog>
      <ErrorDialog
        error={cancelSubscriptionError}
        rules={[]}
        fallbackErrorMessageID="SubscriptionScreen.cancel.error"
      />
      <Dialog
        hidden={enterpriseDialogHidden}
        onDismiss={onDismiss}
        dialogContentProps={enterpriseDialogContentProps}
      >
        <DialogFooter>
          <PrimaryButton
            href="mailto:hello@authgear.com"
            onClick={onDismiss}
            text={<FormattedMessage id="SubscriptionScreen.enterprise.cta" />}
          />
        </DialogFooter>
      </Dialog>

      <div className={styles.root}>
        <div className={cn(styles.section, "grid gap-4 grid-flow-row")}>
          <ScreenTitle>
            <FormattedMessage id="SubscriptionScreen.title" />
          </ScreenTitle>
          <ScreenDescription>
            <FormattedMessage id="SubscriptionScreen.description" />
          </ScreenDescription>
        </div>
        <Pivot
          className="mb-6"
          onLinkClick={onTabChange}
          selectedKey={selectedTab}
        >
          <PivotItem
            itemKey={Tab.Subscription}
            headerText={renderToString("SubscriptionScreen.tabs.subscription")}
          />
          <PivotItem
            itemKey={Tab.PlanDetail}
            headerText={renderToString("SubscriptionScreen.tabs.planDetails")}
          />
        </Pivot>
        {selectedTab === Tab.Subscription ? (
          <>
            <FeatureBanner />
            <PlansSection
              currentPlanName={planName}
              subscriptionCancelled={subscriptionCancelled}
              nextBillingDate={nextBillingDate}
              subscriptionPlans={subscriptionPlans}
              onClickContactUs={onClickContactUs}
              onClickCancelSubscription={onClickCancel}
            />
            <footer className={styles.section}>
              <Text block={true}>
                <FormattedMessage id="SubscriptionScreen.footer.tax" />
              </Text>
            </footer>
          </>
        ) : (
          <></>
        )}
      </div>
    </>
  );
}

interface SubscriptionProcessingPaymentScreenProps {
  stripeError?: stripe.Error;
}

const SubscriptionProcessingPaymentScreen: React.VFC<SubscriptionProcessingPaymentScreenProps> =
  function SubscriptionProcessingPaymentScreen(
    props: SubscriptionProcessingPaymentScreenProps
  ) {
    const { stripeError } = props;
    const { renderToString } = useContext(Context);
    const { appID } = useParams() as { appID: string };

    const {
      cancelFailedSubscription,
      loading: cancelFailedSubscriptionLoading,
      error: cancelFailedSubscriptionError,
    } = useCancelFailedSubscriptionMutation(appID);
    useLoading(cancelFailedSubscriptionLoading);

    const paymentStatus = useMemo(() => {
      if (stripeError == null) {
        return "IsProcessing";
      }
      // https://stripe.com/docs/error-codes
      if (stripeError.code === "card_declined") {
        return "CardDeclined";
      }
      return "UnknownError";
    }, [stripeError]);

    const onClickCancelFailedSubscription = useCallback(async () => {
      await cancelFailedSubscription();
    }, [cancelFailedSubscription]);

    return (
      <div className={styles.root}>
        <ScreenTitle className={styles.section}>
          <FormattedMessage id="SubscriptionScreen.title" />
        </ScreenTitle>
        <div
          className={cn(styles.processingPaymentSection)}
          style={{
            boxShadow: DefaultEffects.elevation4,
          }}
        >
          {paymentStatus === "IsProcessing" ? (
            <>
              <Spinner
                className={styles.processingPaymentSpinner}
                labelPosition="right"
                label={renderToString("SubscriptionScreen.processing-payment")}
                size={SpinnerSize.large}
                styles={{
                  label: {
                    whiteSpace: "pre-line",
                    textAlign: "left",
                    marginLeft: "16px",
                    fontSize: "14px",
                    lineHeight: "20px",
                  },
                }}
              />
              <ThemeProvider theme={CONTACT_US_BUTTON_THEME}>
                <ExternalLink href={"mailto:hello@authgear.com"}>
                  <Text>
                    <FormattedMessage id="SubscriptionScreen.contact-us.label" />
                  </Text>
                </ExternalLink>
              </ThemeProvider>
            </>
          ) : null}
          {paymentStatus === "CardDeclined" ? (
            <>
              <Text className={styles.processingPaymentErrorMessage}>
                <FormattedMessage id="SubscriptionScreen.payment-declined.description" />
              </Text>
              <div className={styles.processingPaymentButtonContainer}>
                <ButtonWithLoading
                  loading={cancelFailedSubscriptionLoading}
                  onClick={onClickCancelFailedSubscription}
                  labelId="SubscriptionScreen.cancel-transaction.label"
                />
              </div>
            </>
          ) : null}
          {paymentStatus === "UnknownError" ? (
            <>
              <Text className={styles.processingPaymentErrorMessage}>
                <FormattedMessage id="SubscriptionScreen.unknown-error.description" />
              </Text>
            </>
          ) : null}
          <ErrorDialog
            error={cancelFailedSubscriptionError}
            rules={[]}
            fallbackErrorMessageID="SubscriptionScreen.cancel-transaction-error.description"
          />
        </div>
      </div>
    );
  };

const SubscriptionScreen: React.VFC = function SubscriptionScreen() {
  const now = useConst(new Date());
  const thisMonth = useMemo(() => {
    return now.toISOString();
  }, [now]);
  const previousMonth = useMemo(() => {
    return DateTime.fromJSDate(now)
      .minus({
        months: 1,
      })
      .toJSDate()
      .toISOString();
  }, [now]);

  const { appID } = useParams() as { appID: string };

  const subscriptionScreenQuery = useSubscriptionScreenQueryQuery({
    variables: {
      id: appID,
      thisMonth,
      previousMonth,
    },
  });

  const isProcessingSubscription =
    !!subscriptionScreenQuery.data &&
    (subscriptionScreenQuery.data.node as AppFragmentFragment)
      .isProcessingSubscription;

  const lastStripeError = useMemo(() => {
    return (
      !!subscriptionScreenQuery.data &&
      (subscriptionScreenQuery.data.node as AppFragmentFragment).lastStripeError
    );
  }, [subscriptionScreenQuery]);

  // if isProcessingSubscription is true
  // refetch in every few seconds and wait until it changes to false
  useEffect(() => {
    if (subscriptionScreenQuery.loading) {
      return () => {};
    }
    if (!isProcessingSubscription) {
      return () => {};
    }
    const interval = setInterval(() => {
      subscriptionScreenQuery.refetch().finally(() => {});
    }, CHECK_IS_PROCESSING_SUBSCRIPTION_INTERVAL);
    return () => {
      clearInterval(interval);
    };
  }, [
    subscriptionScreenQuery.loading,
    isProcessingSubscription,
    subscriptionScreenQuery,
  ]);

  if (subscriptionScreenQuery.loading) {
    return <ShowLoading />;
  }

  if (subscriptionScreenQuery.error) {
    return (
      <ShowError
        error={subscriptionScreenQuery.error}
        onRetry={() => {
          subscriptionScreenQuery.refetch().finally(() => {});
        }}
      />
    );
  }

  if (isProcessingSubscription) {
    return (
      <SubscriptionProcessingPaymentScreen stripeError={lastStripeError} />
    );
  }

  const planName = (subscriptionScreenQuery.data?.node as AppFragmentFragment)
    .planName;
  const subscription = (
    subscriptionScreenQuery.data?.node as AppFragmentFragment
  ).subscription;
  const subscriptionPlans =
    subscriptionScreenQuery.data?.subscriptionPlans ?? [];
  const thisMonthUsage = (
    subscriptionScreenQuery.data?.node as AppFragmentFragment
  ).thisMonth;
  const previousMonthUsage = (
    subscriptionScreenQuery.data?.node as AppFragmentFragment
  ).previousMonth;

  const effectiveAppConfig = (
    subscriptionScreenQuery.data?.node as AppFragmentFragment
  ).effectiveAppConfig as PortalAPIAppConfig | null | undefined;

  return (
    <ScreenLayoutScrollView>
      <SubscriptionScreenContent
        appID={appID}
        planName={planName}
        subscription={subscription ?? undefined}
        subscriptionPlans={subscriptionPlans}
        thisMonthUsage={thisMonthUsage ?? undefined}
        previousMonthUsage={previousMonthUsage ?? undefined}
        effectiveAppConfig={effectiveAppConfig ?? undefined}
      />
    </ScreenLayoutScrollView>
  );
};

export default SubscriptionScreen;
