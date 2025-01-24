import React, { useCallback, useMemo } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { FormattedMessage } from "@oursky/react-messageformat";
import { produce } from "immer";
import cn from "classnames";
import { Text } from "@fluentui/react";
import FormContainer from "../../FormContainer";
import {
  AppConfigFormModel,
  useAppConfigForm,
} from "../../hook/useAppConfigForm";
import ScreenContent from "../../ScreenContent";
import ShowError from "../../ShowError";
import ShowLoading from "../../ShowLoading";
import UserProfileAttributesList, {
  ItemComponentProps,
} from "../../UserProfileAttributesList";
import {
  PortalAPIAppConfig,
  CustomAttributesAttributeConfig,
} from "../../types";
import { parseJSONPointer } from "../../util/jsonpointer";
import styles from "./CustomAttributesConfigurationScreen.module.css";
import PrimaryButton from "../../PrimaryButton";
import NavBreadcrumb from "../../NavBreadcrumb";

interface FormState {
  items: CustomAttributesAttributeConfig[];
}

interface CustomAttributesConfigurationScreenContentProps {
  form: AppConfigFormModel<FormState>;
}

function constructFormState(config: PortalAPIAppConfig): FormState {
  const items = config.user_profile?.custom_attributes?.attributes ?? [];
  return {
    items,
  };
}

function constructConfig(
  rawConfig: PortalAPIAppConfig,
  _initialState: FormState,
  currentState: FormState,
  _effectiveConfig: PortalAPIAppConfig
): PortalAPIAppConfig {
  return produce(rawConfig, (rawConfig) => {
    rawConfig.user_profile ??= {};
    rawConfig.user_profile.custom_attributes ??= {};
    rawConfig.user_profile.custom_attributes.attributes = currentState.items;
  });
}

function EmptyState() {
  const navigate = useNavigate();
  const onClick = useCallback(
    (e: React.MouseEvent<unknown>) => {
      e.preventDefault();
      e.stopPropagation();
      navigate("./add");
    },
    [navigate]
  );
  return (
    <div className={styles.emptyState}>
      <Text className={styles.emptyStateMessage} block={true}>
        <FormattedMessage id="CustomAttributesConfigurationScreen.empty-message" />
      </Text>
      <PrimaryButton
        className={styles.addNewAttributeButton}
        onClick={onClick}
        text={
          <FormattedMessage id="CustomAttributesConfigurationScreen.label.add-new-attribute" />
        }
      />
    </div>
  );
}

function ItemComponent(
  props: ItemComponentProps<CustomAttributesAttributeConfig>
) {
  const { className, item } = props;
  const { pointer } = item;
  const fieldName = parseJSONPointer(pointer)[0];
  return (
    <Text className={cn(className, styles.fieldName)} block={true}>
      {fieldName}
    </Text>
  );
}

const CustomAttributesConfigurationScreenContent: React.VFC<CustomAttributesConfigurationScreenContentProps> =
  function CustomAttributesConfigurationScreenContent(props) {
    const navigate = useNavigate();
    const { state, setState } = props.form;
    const { items } = state;

    const navBreadcrumbItems = useMemo(() => {
      return [
        {
          to: ".",
          label: (
            <FormattedMessage id="CustomAttributesConfigurationScreen.title" />
          ),
        },
      ];
    }, []);

    const onChangeItems = useCallback(
      (newItems: CustomAttributesAttributeConfig[]) => {
        setState((prev) => {
          return {
            ...prev,
            items: newItems,
          };
        });
      },
      [setState]
    );

    const onEditButtonClick = useCallback(
      (index: number) => {
        navigate(`./${index}/edit`);
      },
      [navigate]
    );

    return (
      <>
        <ScreenContent layout="list">
          <div className={styles.widget}>
            <div className="flex gap-x-1">
              <NavBreadcrumb
                className="flex-1 overflow-hidden items-center"
                items={navBreadcrumbItems}
              />
              <PrimaryButton
                text={
                  <FormattedMessage id="CustomAttributesConfigurationScreen.label.add-new-attribute" />
                }
                iconProps={useMemo(() => ({ iconName: "Add" }), [])}
                onClick={useCallback(() => navigate("./add"), [navigate])}
              />
            </div>
          </div>
          <div className={styles.widget}>
            <UserProfileAttributesList
              items={items}
              onChangeItems={onChangeItems}
              onReorderItems={onChangeItems}
              onEditButtonClick={onEditButtonClick}
              ItemComponent={ItemComponent}
            />
            {state.items.length <= 0 ? <EmptyState /> : null}
          </div>
        </ScreenContent>
      </>
    );
  };

const CustomAttributesConfigurationScreen: React.VFC =
  function CustomAttributesConfigurationScreen() {
    const { appID } = useParams() as { appID: string };
    const form = useAppConfigForm({
      appID,
      constructFormState,
      constructConfig,
    });

    if (form.isLoading) {
      return <ShowLoading />;
    }

    if (form.loadError) {
      return <ShowError error={form.loadError} onRetry={form.reload} />;
    }

    return (
      <FormContainer
        form={form}
        showDiscardButton={true}
        stickyFooterComponent={true}
      >
        <CustomAttributesConfigurationScreenContent form={form} />
      </FormContainer>
    );
  };

export default CustomAttributesConfigurationScreen;
