import { Injectable, Inject, DOCUMENT } from '@angular/core'
import { Subject, Observable, OperatorFunction, pipe, interval } from 'rxjs'
import { WebSocketSubject, webSocket } from 'rxjs/webSocket'
import { filter, delay, tap, retry, delayWhen, finalize } from 'rxjs/operators'

export const enum SocketAction {
    Register = 'register',
    Reload = 'reload',
    Message = 'message',
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export interface SocketMessage<T = any> {
    uuid?: string
    action: SocketAction
    payload: T
}

@Injectable({
    providedIn: 'root',
})
export class SocketService {
    private readonly url: string = `ws://${this.document.location.host}/socket`

    private readonly socket$: WebSocketSubject<SocketMessage> = webSocket(this.url)

    private readonly subject$: Subject<SocketMessage> = new Subject()

    public constructor(@Inject(DOCUMENT) private readonly document: Document) {
        this.socket$.pipe(this.reconnect(5000)).subscribe()
    }

    /**
     * Subscribe for {@link SocketAction}s from the server.
     *
     * @param action
     * @returns
     */
    public subscribe<T = string>(action: SocketAction): Observable<SocketMessage<T>> {
        return this.subject$.pipe(
            filter((message: SocketMessage<T>) => message?.action === action),
            delay(500),
        )
    }

    /**
     * Send {@link SocketAction}s to the server.
     *
     * @param action
     * @param payload
     */
    public next<A extends SocketAction, P>(action: A, payload?: P): void {
        this.socket$.next({ action, payload })
    }

    /**
     * Catch socket closing events and automatically try to reconnect.
     *
     * @param after
     * @returns
     */
    private reconnect<T extends SocketMessage>(after: number = 5000): OperatorFunction<T, T> {
        return pipe(
            tap(() => console.info('[socket] Connection established.')),
            retry({
                delay: error => {
                    return error.pipe(
                        tap(({ type }: Event): void => {
                            if (type === 'close') {
                                console.info('[socket] Disconnected!')
                            }

                            console.info('[socket] Trying to reconnect...')
                        }),
                        delayWhen(() => interval(after)),
                    )
                },
            }),
            tap(message => this.subject$.next(message)),
            finalize(() => this.subject$.complete()),
        )
    }
}
