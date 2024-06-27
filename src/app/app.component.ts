import { Component, OnInit, DestroyRef, inject } from '@angular/core'
import { takeUntilDestroyed } from '@angular/core/rxjs-interop'
import { Router, RouterOutlet } from '@angular/router'
import { Store } from '@ngrx/store'
import { NgProgressModule } from 'ngx-progressbar'
import { SocketService, SocketAction } from 'shared/socket/socket.service'
import { HealthService, Status } from 'shared/health/health.service'
import { ReloadService } from 'shared/reload/reload.service'
import { SettingsActions } from 'store/settings'

@Component({
    standalone: true,
    selector: 'cmp-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
    imports: [RouterOutlet, NgProgressModule],
})
export class AppComponent implements OnInit {
    private readonly destroy: DestroyRef = inject(DestroyRef)

    public constructor(
        private readonly router: Router,
        private readonly store: Store,
        private readonly socket: SocketService,
        private readonly health: HealthService,
        private readonly reload: ReloadService,
    ) {}

    public ngOnInit(): void {
        // delayed initial navigation
        setTimeout(() => this.router.initialNavigation())

        // fetch settings from server
        this.store.dispatch(SettingsActions.load())

        // watch for reload actions
        this.socket
            .subscribe(SocketAction.Reload)
            .pipe(takeUntilDestroyed(this.destroy))
            .subscribe(() => {
                this.reload.reload()
            })

        // watch for unhealthy state
        this.health
            .watch(Status.Unavailable)
            .pipe(takeUntilDestroyed(this.destroy))
            .subscribe(() => {
                this.reload.reload()
            })
    }
}
