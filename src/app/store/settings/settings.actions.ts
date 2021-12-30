import { createAction, props } from '@ngrx/store'
import { Settings } from './settings.model'

export enum SettingsActions {
    LOAD = '[Settings] Load settings',
    LOAD_SUCCESS = '[Settings] Load settings success',
    LOAD_FAIL = '[Settings] Load settings fail',
}

/**
 * Load settings
 */
export const LoadSettings = createAction(SettingsActions.LOAD)

export const LoadSettingsSuccess = createAction(
    SettingsActions.LOAD_SUCCESS,
    props<{ payload: Settings }>(),
)

export const LoadSettingsFail = createAction(
    SettingsActions.LOAD_FAIL,
    props<{ payload: string; error?: Error }>(),
)
