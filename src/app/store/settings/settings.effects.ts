import { Injectable } from '@angular/core'
import { createEffect, Actions, ofType } from '@ngrx/effects'
import { of } from 'rxjs'
import { mergeMap, map, catchError } from 'rxjs/operators'
import { SocketService, SocketAction } from 'shared/socket/socket.service'
import { SettingsActions, LoadSettingsSuccess, LoadSettingsFail } from './settings.actions'
import { Settings } from './settings.model'

@Injectable()
export class SettingsEffects {
    public getSettings$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(SettingsActions.LOAD),
            mergeMap(() => {
                return this.socket.subscribe<Settings>(SocketAction.Register).pipe(
                    map(({ payload }) => LoadSettingsSuccess({ payload })),
                    catchError(error => {
                        return of(LoadSettingsFail({ error, payload: 'SETTINGS_LOAD_ERROR' }))
                    }),
                )
            }),
        )
    })

    public constructor(
        private readonly actions$: Actions,
        private readonly socket: SocketService,
    ) {}
}
