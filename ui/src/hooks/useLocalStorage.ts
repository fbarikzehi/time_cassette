import { useState } from "react";


export default function useLocalStorage() {

    enum localStorageKeys {
        AppState = 'app_state',
    };

    const [data, setData] = useState()

    const getStorage = (key: localStorageKeys) => {
        try {
            const value = localStorage.getItem(key);
            if (!value) return undefined;
            setData(JSON.parse(value));
        } catch (e) {
            return undefined;
        }
    }

    const setStorage = (key: localStorageKeys) => {
        try {
            const value = JSON.stringify(data);
            localStorage.setItem(key, value);
        } catch (e) {

        }
    }


    return { data, setData, getStorage, setStorage, localStorageKeys }

}
