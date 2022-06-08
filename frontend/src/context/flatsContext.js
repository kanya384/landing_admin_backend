import {createContext} from 'react'
const noop = ()=> {}
export const FlatsContext = createContext ({
    flats: [],
    activeFlat:null,
    setFlats:noop(),
    setActiveFlat:noop(),
})
