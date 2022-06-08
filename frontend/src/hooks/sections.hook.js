import {useState} from 'react'

export const useSections = () => {
    const[blocks, setBlocks_] = useState(3)
    const[menuClick, setMenuClick] = useState(false)
    const setBlocks = (count) =>{
        setBlocks_(count)
    }
    return {blocks, setBlocks, menuClick, setMenuClick}
}
