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

// 初期値（string型である点に注意）
const initialUserIdString: string = localStorage.getItem("userId") ?? "-1";
const initialTokenString: string = localStorage.getItem("tokenString") ?? "";

const initialUserId: number = Math.round(Number(initialUserIdString));

const initialAuth: Auth = {
  userId: initialUserId,
  tokenString: initialTokenString,
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
      localStorage.setItem("userId", action.payload.userId.toString());
      localStorage.setItem("tokenString", action.payload.tokenString);
      return {
        userId: action.payload.userId,
        tokenString: action.payload.tokenString,
      };
    case clearAuthType:
      console.log("clear auth reducer");
      localStorage.setItem("userId", initialAuth.userId.toString());
      localStorage.setItem("tokenString", initialAuth.tokenString);
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
