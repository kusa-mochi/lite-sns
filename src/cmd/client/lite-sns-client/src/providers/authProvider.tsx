import { createContext, Dispatch, ReactNode, useContext, useEffect, useState } from "react"

export type Auth = {
    userId: number
    tokenString: string
}

const AuthContext = createContext<Auth | null>(null)
const SetAuthContext = createContext<Dispatch<Auth | null> | null>(null)

export function useAuth(): Auth {
    const auth = useContext(AuthContext)
    if (!auth) throw new Error("wrap this component by AuthProvider")
    
    return auth
}

export function useSetAuth(): Dispatch<Auth | null> {
    const setAuth = useContext(SetAuthContext)
    if (!setAuth) throw new Error("wrap this component by AuthProvider")
    
    return setAuth
}

type AuthProviderProps = {
    children: ReactNode
}
export const AuthProvider = (props: AuthProviderProps) => {
    const [auth, setAuth] = useState<Auth | null>({
        userId: -1,
        tokenString: "",
    })

    if (!auth) return <div>Loading...</div>

    return (
        <AuthContext.Provider value={auth}>
            <SetAuthContext.Provider value={setAuth}>
                {props.children}
            </SetAuthContext.Provider>
        </AuthContext.Provider>
    )
}
