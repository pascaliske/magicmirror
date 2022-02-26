import { Component, OnInit } from '@angular/core'
import { UntilDestroy, untilDestroyed } from '@ngneat/until-destroy'
import { Store } from '@ngrx/store'
import { SocketService, SocketAction } from 'shared/socket/socket.service'
import { HealthService, Status } from 'shared/health/health.service'
import { ReloadService } from 'shared/reload/reload.service'
import { AppState } from 'store'
import { LoadSettings } from 'store/settings'

@UntilDestroy()
@Component({
    selector: 'cmp-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
    public constructor(
        private readonly store: Store<AppState>,
        private readonly socket: SocketService,
        private readonly health: HealthService,
        private readonly reload: ReloadService,
    ) {}

    public ngOnInit(): void {
        // fetch settings from server
        this.store.dispatch(LoadSettings())

        // watch for reload actions
        this.socket
            .subscribe(SocketAction.Reload)
            .pipe(untilDestroyed(this))
            .subscribe(() => {
                this.reload.reload()
            })

        // watch for unhealthy state
        this.health
            .watch(Status.Unavailable)
            .pipe(untilDestroyed(this))
            .subscribe(() => {
                this.reload.reload()
            })
    }
}
