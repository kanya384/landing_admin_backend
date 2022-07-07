import {useContext, useEffect, useState} from 'react'
import { ContentContext } from '../context/contentContext'

export const useSections = () => {
    const[blocks, setBlocks_] = useState(window.location.search.indexOf('liter')!==-1?30:3)
    const[menuClick, setMenuClick] = useState(false)
    const[showPlan, setShowPlan] = useState(false)
    const setBlocks = (count) =>{
        setBlocks_(count)
    }
    
    return {blocks, setBlocks, menuClick, setMenuClick, showPlan, setShowPlan}
}
