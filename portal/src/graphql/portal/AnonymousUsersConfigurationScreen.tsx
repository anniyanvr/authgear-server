import React, { useCallback, useContext, useMemo } from "react";
import { useParams } from "react-router-dom";
import { produce } from "immer";
import { Context, FormattedMessage } from "@oursky/react-messageformat";
import {
  Dropdown,
  IDropdownOption,
  Toggle,
  Text,
  DetailsList,
  SelectionMode,
  IColumn,
  IDetailsHeaderProps,
  DetailsHeader,
} from "@fluentui/react";
import {
  isPromotionConflictBehaviour,
  PortalAPIAppConfig,
  PromotionConflictBehaviour,
  promotionConflictBehaviours,
  OAuthClientConfig,
} from "../../types";
import { clearEmptyObject } from "../../util/misc";
import ShowLoading from "../../ShowLoading";
import ShowError from "../../ShowError";
import ScreenContent from "../../ScreenContent";
import ScreenTitle from "../../ScreenTitle";
import ScreenDescription from "../../ScreenDescription";
import WidgetTitle from "../../WidgetTitle";
import Widget from "../../Widget";
import {
  AppConfigFormModel,
  useAppConfigForm,
} from "../../hook/useAppConfigForm";
import FormContainer from "../../FormContainer";
import styles from "./AnonymousUsersConfigurationScreen.module.scss";

const dropDownStyles = {
  dropdown: {
    width: "300px",
  },
};

interface FormState {
  enabled: boolean;
  promotionConflictBehaviour: PromotionConflictBehaviour;
  oauthClients: OAuthClientConfig[];
  sessionPersistentCookie: boolean;
  sessionLifetimeSeconds: number | undefined;
  sessionIdleTimeoutEnabled: boolean;
  sessionIdleTimeoutSeconds: number | undefined;
}

function constructFormState(config: PortalAPIAppConfig): FormState {
  const enabled =
    config.authentication?.identities?.includes("anonymous") ?? false;
  const promotionConflictBehaviour =
    config.identity?.on_conflict?.promotion ?? "error";
  const oauthClients = config.oauth?.clients ?? [];
  return {
    enabled,
    promotionConflictBehaviour,
    oauthClients,
    sessionPersistentCookie: !(config.session?.cookie_non_persistent ?? false),
    sessionLifetimeSeconds: config.session?.lifetime_seconds,
    sessionIdleTimeoutEnabled: config.session?.idle_timeout_enabled ?? false,
    sessionIdleTimeoutSeconds: config.session?.idle_timeout_seconds,
  };
}

function constructConfig(
  config: PortalAPIAppConfig,
  initialState: FormState,
  currentState: FormState,
  effectiveConfig: PortalAPIAppConfig
): PortalAPIAppConfig {
  // eslint-disable-next-line complexity
  return produce(config, (config) => {
    if (initialState.enabled !== currentState.enabled) {
      const identities = (
        effectiveConfig.authentication?.identities ?? []
      ).slice();
      const index = identities.indexOf("anonymous");
      if (currentState.enabled && index === -1) {
        identities.push("anonymous");
      } else if (!currentState.enabled && index >= 0) {
        identities.splice(index, 1);
      }
      config.authentication ??= {};
      config.authentication.identities = identities;
    }
    if (
      currentState.enabled &&
      initialState.promotionConflictBehaviour !==
        currentState.promotionConflictBehaviour
    ) {
      config.identity ??= {};
      config.identity.on_conflict ??= {};
      config.identity.on_conflict.promotion =
        currentState.promotionConflictBehaviour;
    }
    clearEmptyObject(config);
  });
}

const conflictBehaviourMessageId: Record<PromotionConflictBehaviour, string> = {
  login: "AnonymousIdentityConflictBehaviour.login",
  error: "AnonymousIdentityConflictBehaviour.error",
};

interface OAuthClientListItem {
  name: string;
  refreshTokenIdleTimeout: string;
  refreshTokenLifetime: string;
}

interface AnonymousUserLifeTimeDescriptionProps {
  form: AppConfigFormModel<FormState>;
}

const AnonymousUserLifeTimeDescription: React.FC<AnonymousUserLifeTimeDescriptionProps> =
  function AnonymousUserLifeTimeDescription(props) {
    const { renderToString } = useContext(Context);
    const {
      sessionIdleTimeoutEnabled,
      sessionIdleTimeoutSeconds,
      sessionLifetimeSeconds,
      sessionPersistentCookie,
      oauthClients,
    } = props.form.state;

    const columns: IColumn[] = useMemo(
      () => [
        {
          key: "name",
          name: renderToString(
            "AnonymousUsersConfigurationScreen.user-lifetime.token.application-name.label"
          ),
          minWidth: 150,
          maxWidth: 150,
          isMultiline: true,
        },
        {
          key: "refreshTokenIdleTimeout",
          name: renderToString(
            "AnonymousUsersConfigurationScreen.user-lifetime.token.refresh-token-idle-timeout.label"
          ),
          minWidth: 150,
          maxWidth: 150,
        },
        {
          key: "refreshTokenLifetime",
          name: renderToString(
            "AnonymousUsersConfigurationScreen.user-lifetime.token.refresh-token-lifetime.label"
          ),
          minWidth: 150,
          maxWidth: 150,
        },
      ],
      [renderToString]
    );

    const items: OAuthClientListItem[] = useMemo(() => {
      return oauthClients.map((client) => {
        return {
          name: client.name ?? "",
          refreshTokenIdleTimeout: client.refresh_token_idle_timeout_enabled
            ? client.refresh_token_idle_timeout_seconds?.toFixed(0) ?? ""
            : "-",
          refreshTokenLifetime:
            client.refresh_token_lifetime_seconds?.toFixed(0) ?? "",
        };
      });
    }, [oauthClients]);

    const onRenderItemColumn = useCallback(
      (item?: OAuthClientListItem, _index?: number, column?: IColumn) => {
        if (item == null) {
          return null;
        }
        switch (column?.key) {
          case "name":
            return item.name;
          case "refreshTokenIdleTimeout":
            return item.refreshTokenIdleTimeout;
          case "refreshTokenLifetime":
            return item.refreshTokenLifetime;
          default:
            return null;
        }
      },
      []
    );

    const onRenderDetailsHeader = useCallback((props?: IDetailsHeaderProps) => {
      if (props == null) {
        return null;
      }
      return <DetailsHeader {...props} className={styles.detailsHeader} />;
    }, []);

    return (
      <Widget className={styles.widget}>
        <WidgetTitle>
          <FormattedMessage id="AnonymousUsersConfigurationScreen.user-lifetime.title" />
        </WidgetTitle>
        <Text variant="medium" block={true}>
          <FormattedMessage id="AnonymousUsersConfigurationScreen.user-lifetime.description" />
        </Text>
        <div>
          <Text className={styles.title} variant="medium" block={true}>
            <FormattedMessage id="AnonymousUsersConfigurationScreen.user-lifetime.cookie.title" />
          </Text>
          {sessionIdleTimeoutEnabled && (
            <Text variant="medium" block={true} className={styles.sessionItem}>
              <FormattedMessage
                id="AnonymousUsersConfigurationScreen.user-lifetime.cookie.idle-timeout.description"
                values={{
                  seconds: sessionIdleTimeoutSeconds?.toFixed(0) ?? "",
                }}
              />
            </Text>
          )}
          <Text variant="medium" block={true} className={styles.sessionItem}>
            <FormattedMessage
              id="AnonymousUsersConfigurationScreen.user-lifetime.cookie.session-lifetime.description"
              values={{
                seconds: sessionLifetimeSeconds?.toFixed(0) ?? "",
              }}
            />
          </Text>
          <Text variant="medium" block={true} className={styles.sessionItem}>
            <FormattedMessage id="AnonymousUsersConfigurationScreen.user-lifetime.cookie.persistent-cookie.label" />{" "}
            :{" "}
            <FormattedMessage
              id={sessionPersistentCookie ? "enabled" : "disabled"}
            />
          </Text>
        </div>
        <div>
          <Text className={styles.title} variant="medium" block={true}>
            <FormattedMessage id="AnonymousUsersConfigurationScreen.user-lifetime.token.title" />
          </Text>
          <DetailsList
            columns={columns}
            items={items}
            selectionMode={SelectionMode.none}
            onRenderItemColumn={onRenderItemColumn}
            onRenderDetailsHeader={onRenderDetailsHeader}
          />
        </div>
        <Text variant="medium" block={true}>
          <FormattedMessage
            id="AnonymousUsersConfigurationScreen.user-lifetime.go-to-applications.description"
            values={{
              applicationsPath: "../apps",
            }}
          />
        </Text>
      </Widget>
    );
  };

interface AnonymousUserConfigurationContentProps {
  form: AppConfigFormModel<FormState>;
}

const AnonymousUserConfigurationContent: React.FC<AnonymousUserConfigurationContentProps> =
  function AnonymousUserConfigurationContent(props) {
    const { state, setState } = props.form;

    const { renderToString } = useContext(Context);

    const conflictBehaviourOptions = useMemo(
      () =>
        promotionConflictBehaviours.map((behaviour) => {
          const selectedBehaviour = state.promotionConflictBehaviour;
          return {
            key: behaviour,
            text: renderToString(conflictBehaviourMessageId[behaviour]),
            isSelected: selectedBehaviour === behaviour,
          };
        }),
      [state, renderToString]
    );

    const onEnableChange = useCallback(
      (_event, checked?: boolean) =>
        setState((state) => ({
          ...state,
          enabled: checked ?? false,
        })),
      [setState]
    );

    const onConflictOptionChange = useCallback(
      (_event, option?: IDropdownOption) => {
        const key = option?.key;
        if (key && isPromotionConflictBehaviour(key)) {
          setState((state) => ({
            ...state,
            promotionConflictBehaviour: key,
          }));
        }
      },
      [setState]
    );

    return (
      <ScreenContent>
        <ScreenTitle className={styles.widget}>
          <FormattedMessage id="AnonymousUsersConfigurationScreen.title" />
        </ScreenTitle>
        <ScreenDescription className={styles.widget}>
          <FormattedMessage id="AnonymousUsersConfigurationScreen.description" />
        </ScreenDescription>
        <Widget className={styles.widget}>
          <WidgetTitle>
            <FormattedMessage id="AnonymousUsersConfigurationScreen.title" />
          </WidgetTitle>
          <Toggle
            checked={state.enabled}
            onChange={onEnableChange}
            label={renderToString(
              "AnonymousUsersConfigurationScreen.enable.label"
            )}
            inlineLabel={true}
          />
          <Dropdown
            styles={dropDownStyles}
            label={renderToString(
              "AnonymousUsersConfigurationScreen.conflict-droplist.label"
            )}
            disabled={!state.enabled}
            options={conflictBehaviourOptions}
            selectedKey={state.promotionConflictBehaviour}
            onChange={onConflictOptionChange}
          />
        </Widget>
        <AnonymousUserLifeTimeDescription form={props.form} />
      </ScreenContent>
    );
  };

const AnonymousUserConfigurationScreen: React.FC =
  function AnonymousUserConfigurationScreen() {
    const { appID } = useParams();
    const form = useAppConfigForm(appID, constructFormState, constructConfig);

    if (form.isLoading) {
      return <ShowLoading />;
    }

    if (form.loadError) {
      return <ShowError error={form.loadError} onRetry={form.reload} />;
    }

    return (
      <FormContainer form={form}>
        <AnonymousUserConfigurationContent form={form} />
      </FormContainer>
    );
  };

export default AnonymousUserConfigurationScreen;
