import {
  createContext,
  Dispatch,
  ReactNode,
  Reducer,
  useContext,
  useReducer,
} from "react";

export const setAuthType = Symbol("SET_AUTH");
export const clearAuthType = Symbol("CLEAR_AUTH");

type ActionType =
  | { type: typeof setAuthType; payload: Auth }
  | { type: typeof clearAuthType; payload: null };

export type Auth = {
  userId: number;
  tokenString: string;
};

const initialAuth: Auth = {
  userId: -1,
  tokenString: "",
};

const AuthContext = createContext<Auth>(initialAuth);
const AuthDispatchContext = createContext<Dispatch<ActionType>>(() => {});

export function useAuth(): Auth {
  const auth = useContext(AuthContext);
  if (!auth) throw new Error("wrap this component by AuthProvider");

  return auth;
}

export function useSetAuth(): Dispatch<ActionType> {
  const setAuth = useContext(AuthDispatchContext);
  if (!setAuth) throw new Error("wrap this component by AuthProvider");

  return setAuth;
}

const authReducer: Reducer<Auth, ActionType> = (
  prevState: Auth,
  action: ActionType
) => {
  switch (action.type) {
    case setAuthType:
      console.log("set auth reducer");
      return {
        userId: action.payload.userId,
        tokenString: action.payload.tokenString,
      };
    case clearAuthType:
      console.log("clear auth reducer");
      return initialAuth;
    default:
      console.log("default auth reducer");
      return prevState;
  }
};

type AuthProviderProps = {
  children: ReactNode;
};
export const AuthProvider = (props: AuthProviderProps) => {
  const [auth, dispatch] = useReducer(authReducer, initialAuth);

  if (!auth) return <div>Loading...</div>;

  return (
    <AuthContext.Provider value={auth}>
      <AuthDispatchContext.Provider value={dispatch}>
        {props.children}
      </AuthDispatchContext.Provider>
    </AuthContext.Provider>
  );
};
