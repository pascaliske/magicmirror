import { SettingsState, settings } from './settings'
import { initialState as settingsInitialState } from './settings/settings.reducer'

export interface AppState {
    settings: SettingsState
}

export const initialState: AppState = {
    settings: settingsInitialState,
}

export const reducers = { settings }
