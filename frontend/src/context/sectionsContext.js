import { createContext } from 'react'
const noop = () => { }
export const SectionsContext = createContext({
    blocks: 3,
    popup: false,
    setPopup: noop(),
    setBlocks: noop(),
    menuClick: false,
    setMenuClick: noop(),
})
