import {useState} from 'react'

export const useContent = () => {
    const[content, setContent] = useState()
    return {content, setContent}
}