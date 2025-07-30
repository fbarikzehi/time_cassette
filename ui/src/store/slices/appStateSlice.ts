import { createSlice, PayloadAction } from "@reduxjs/toolkit";



const initialState = {
  activeMenu: '/workstation'
} as AppStateModel;

export const appStateSlice = createSlice({
  name: "app_state",
  initialState,
  reducers: {
    updateAppState: (state, action: PayloadAction<AppStateModel>) => {
      return state = { ...state, ...action.payload };
    },
  },
});

export const { updateAppState } = appStateSlice.actions;
export default appStateSlice.reducer;