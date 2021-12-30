import { trigger, transition, style, animate } from '@angular/animations'

const visible = style({
    opacity: 1,
})

const hidden = style({
    opacity: 0,
})

export const animations = [
    trigger('fade', [
        transition(':enter', [hidden, animate('1.5s ease-out', visible)]),
        transition(':leave', [animate('1.5s ease-out', hidden)]),
    ]),
]
