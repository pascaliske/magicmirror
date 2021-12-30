import { createReducer, on } from '@ngrx/store'
import { Settings } from './settings.model'
import { LoadSettings, LoadSettingsSuccess, LoadSettingsFail } from './settings.actions'

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

export const settings = createReducer(
    initialState,
    on(LoadSettings, state => ({ ...state, loading: true, loaded: false })),
    on(LoadSettingsSuccess, (state, { payload }) => ({
        ...state,
        data: payload,
        loading: false,
        loaded: true,
    })),
    on(LoadSettingsFail, (state, { payload }) => ({
        ...state,
        error: payload,
        loading: false,
        loaded: false,
    })),
)
