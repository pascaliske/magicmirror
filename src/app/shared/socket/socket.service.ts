import { Injectable, Inject } from '@angular/core'
import { DOCUMENT } from '@angular/common'
import { Observable, EMPTY } from 'rxjs'
import { filter, delay, catchError } from 'rxjs/operators'
import { webSocket, WebSocketSubject } from 'rxjs/webSocket'

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

    public constructor(@Inject(DOCUMENT) private readonly document: Document) {}

    public subscribe<T = string>(action: SocketAction): Observable<SocketMessage<T>> {
        return this.socket$.pipe(
            filter((message: SocketMessage<T>) => message.action === action),
            delay(500),
            catchError((error: Error) => {
                // eslint-disable-next-line no-console
                console.log('==> error', error)
                return EMPTY
            }),
        )
    }

    public next<A extends SocketAction, P>(action: A, payload?: P): void {
        this.socket$.next({ action, payload })
    }
}
