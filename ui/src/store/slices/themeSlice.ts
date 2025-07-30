import { createSlice, PayloadAction } from "@reduxjs/toolkit";



const initialState = {
    theme: {
        mode: 'light'
    },
} as ThemeStateModel;

export const themeSlice = createSlice({
    name: "theme",
    initialState,
    reducers: {
        change: (state, action: PayloadAction<ThemeStateModel>) => {
            state.theme = action.payload.theme;
        },
    },
});

export const { change } = themeSlice.actions;
export default themeSlice.reducer;