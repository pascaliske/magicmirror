import { createReducer, on } from '@ngrx/store'
import { Settings } from './settings.model'
import { SettingsActions } from './settings.actions'

export interface SettingsState {
    data: Settings
    error: string | null
    loaded: boolean
    loading: boolean
}

export const initialState: SettingsState = {
    data: {} as Settings,
    error: null,
    loaded: false,
    loading: false,
}

export const settingsReducer = createReducer(
    initialState,

    // load
    on(SettingsActions.load, (state: SettingsState) => ({
        ...state,
        loading: true,
        loaded: false,
    })),
    on(SettingsActions.loadSuccess, (state: SettingsState, { payload }) => ({
        ...state,
        data: payload,
        loading: false,
        loaded: true,
    })),
    on(SettingsActions.loadError, (state: SettingsState, { payload }) => ({
        ...state,
        error: payload,
        loading: false,
        loaded: false,
    })),
)
