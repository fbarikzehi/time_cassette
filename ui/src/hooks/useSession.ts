import { cookies } from 'next/headers'
import { useCallback, useEffect, useState } from 'react';

// token: '',
// expires: '',
// email: '',
// firstName: '',
// lastName: '',
// avatar: '',
// phoneNumber: ''

export function useSession(keepOnWindowClosed = true) {
    const sessionKey = '_session'

    const getStorage = useCallback(() => {
        return keepOnWindowClosed ? localStorage : sessionStorage;
    }, [keepOnWindowClosed]);

    const getStorageValue = useCallback(() => {
        try {
            const storageValue = getStorage().getItem(sessionKey);
            if (storageValue != null) {
                try {
                    const session = JSON.parse(storageValue);
                    return session;
                } catch (_a) {
                    return storageValue;
                }
            }
        } catch (_b) {
            console.warn(
                "useSession could not access the browser storage. Session will be lost when closing browser window"
            );
        }
        return null;
    }, [getStorage, sessionKey]);

    const [session, setState] = useState(getStorageValue);

    const setSession = (sessionValue: any) => {
        if (typeof sessionValue == "object" || typeof sessionValue === "string") {
            getStorage().setItem(sessionKey, JSON.stringify(sessionValue));
            setState(sessionValue);
        } else {
            throw new Error(
                "useSession hook only accepts objects or strings as session values"
            );
        }
    };

    const clearSession = () => {
        getStorage().removeItem(sessionKey);
        setState(null);
    };

    const isValidSession = () => {
        //TODO: security check for session expiration and / return user after resingnin to last url
        setState(getStorageValue());
        return session != undefined || session != null;
    };

    const syncState = useCallback((event: any) => {
        if (event.key === sessionKey) {
            setState(getStorageValue());
        }
    }, [sessionKey, getStorageValue]);


    useEffect(() => {
        window.addEventListener("storage", syncState);
        return () => {
            window.removeEventListener("storage", syncState);
        };
    }, [sessionKey, syncState]);


    return { session, setSession, clearSession, isValidSession, getStorage };
}