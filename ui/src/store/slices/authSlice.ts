import { createSlice, PayloadAction } from "@reduxjs/toolkit";



const initialState = {
  token: '',
  expires: '',
} as AuthStateModel;

export const authStateSlice = createSlice({
  name: "app_state",
  initialState,
  reducers: {
    setToken: (state, action: PayloadAction<AuthStateModel>) => {
      return state = action.payload;
    },
  },
});

export const { setToken } = authStateSlice.actions;
export default authStateSlice.reducer;