import { Injectable, inject } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Observable, interval, of, EMPTY } from 'rxjs'
import { concatMap, filter, map, catchError } from 'rxjs/operators'
import { environment } from 'environments/environment'

export const enum Status {
    Available = 'OK',
    PartiallyAvailable = 'Partially Available',
    Unavailable = 'Unavailable',
    Timeout = 'Timeout during health check',
}

@Injectable({
    providedIn: 'root',
})
export class HealthService {
    private readonly http: HttpClient = inject(HttpClient)

    /**
     * Interval in milliseconds between subsequent health checks.
     */
    private readonly interval: number = 5 * 1000 * 60

    /**
     * Watch for a specific health check status.
     *
     * @param status Health status to watch for. Defaults to {@link Status.Available}
     * @param customInterval Interval in milliseconds between subsequent checks. Defaults to 5 minutes.
     * @returns An Observable for the requested status.
     */
    public watch(status: Status = Status.Available, customInterval?: number): Observable<Status> {
        // health check should be run on production only
        if (!environment.production) {
            return EMPTY
        }

        return interval(customInterval ?? this.interval).pipe(
            concatMap(() => this.performCheck()),
            filter(result => result === status),
        )
    }

    private performCheck(): Observable<Status> {
        return this.http.get<{ status: Status }>(`/health`).pipe(
            map(({ status }) => status),
            catchError(() => of(Status.Unavailable)),
        )
    }
}
