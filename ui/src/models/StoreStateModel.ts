interface ThemeStateModel {
    theme: {
        mode: String
    }
}

interface AppStateModel {
    activeMenu: String
}

interface AuthStateModel {
    token: String,
    expires: String
}

interface storeModel {
    theme: ThemeStateModel,
    appState: AppStateModel
}

