import {useState} from 'react'

export const useContent = () => {
    const[content, setContent] = useState()
    const [administrate, setAdministrate] = useState(false)

    return {content, setContent, administrate, setAdministrate}
}