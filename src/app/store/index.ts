import { isDevMode } from '@angular/core'
import { ActionReducerMap, MetaReducer } from '@ngrx/store'
import { StoreDevtoolsOptions } from '@ngrx/store-devtools'
import { RouterReducerState, routerReducer } from '@ngrx/router-store'
import { SettingsState, settingsReducer } from './settings/settings.reducer'

export interface State {
    router: RouterReducerState
    settings: SettingsState
}

export const reducers: ActionReducerMap<State> = {
    router: routerReducer,
    settings: settingsReducer,
}

export const metaReducers: MetaReducer<State>[] = isDevMode() ? [] : []

export const storeDevtoolsOptions: StoreDevtoolsOptions = { maxAge: 25, logOnly: !isDevMode() }
