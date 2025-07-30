import { combineReducers, configureStore } from "@reduxjs/toolkit";
import themeSlice from "./slices/themeSlice";
import appStateSlice from "./slices/appStateSlice";
import appStateMiddleware from "./middlewares/appStateMiddleware";

const rootReducer = combineReducers({
  theme: themeSlice,
  appState: appStateSlice
},);

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(appStateMiddleware),
});
