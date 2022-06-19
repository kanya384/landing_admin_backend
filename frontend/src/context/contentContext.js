import {createContext} from 'react'
const noop = ()=> {}
export const ContentContext = createContext ({
    content: {},
    setContent:noop(),
    administrate: false,
    setAdministrate: noop(),
})