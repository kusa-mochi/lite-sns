import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useState,
} from "react";
import config from "../conf/release/lite-sns-client.json";

// app server config
// NOTE: この値は export せず、useConfig() の戻り値からアクセスして使う想定。
const configData: LiteSnsClientConfig = {
  appServer: {
    apiPrefix: config.app.api_prefix,
    ip: config.app.ip,
    port: config.app.port,
  },
};

type AppServerInfo = {
  apiPrefix: string;
  ip: string;
  port: number;
};

type LiteSnsClientConfig = {
  appServer: AppServerInfo;
};

const ConfigContext = createContext<LiteSnsClientConfig | null>(null);

// ConfigContextを使うコンポーネントが呼び出すフック
export function useConfig(): LiteSnsClientConfig {
  const theme = useContext(ConfigContext);
  if (!theme) throw new Error("wrap this component by ConfigProvider");

  return theme;
}

type ConfigProviderProps = {
  children: ReactNode;
};
export const ConfigProvider = (props: ConfigProviderProps) => {
  const [config, setConfig] = useState<LiteSnsClientConfig | null>(null);

  useEffect(() => {
    setConfig(configData);
  }, []);

  if (!config) return <div>Loading...</div>;

  return (
    <ConfigContext.Provider value={config}>
      {props.children}
    </ConfigContext.Provider>
  );
};
