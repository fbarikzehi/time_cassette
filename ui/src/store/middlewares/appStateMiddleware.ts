import { updateAppState } from "../slices/appStateSlice";
import { setToken } from "../slices/authSlice";


const appStateMiddleware = (store: any) => (next: any) => (action: any) => {
    if (updateAppState.match(action)) {
        localStorage.setItem("app_state", JSON.stringify(action.payload));
    }
    else if (setToken.match(action)) {
        localStorage.setItem("_session", JSON.stringify(action.payload));
    }
    return next(action);
};



export default appStateMiddleware;
