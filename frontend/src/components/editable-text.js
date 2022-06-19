import { height } from "@mui/system"
import parse from "html-react-parser";
import {Fragment, useContext, useEffect, useState} from "react"
import TextareaAutosize from 'react-textarea-autosize';
import { ContentContext } from "../context/contentContext";

export const EditableText = ({id, defaultText}) => {
    const content = useContext(ContentContext)
    const [text, setText] = useState("")
    const [index, setIndex]= useState()
    const getIndexFromEditable = () =>{
        let index = -1
        if (content.content && content.content.Editables){
            content.content.Editables.forEach((editable, ind)=>{
                if (editable.ID == id) {
                    index= ind
                }
            })
        } 
        return index
    }

    const handleChange = (e) => {
        setText(e.target.value)
        let editableNew = content.content.Editables
        editableNew[index].value = e.target.value
        content.setContent({
            ...content.content,
            Editable:editableNew
        })
    }

    useEffect(()=>{
        if (content.content && content.content.Editables){
            let index = getIndexFromEditable()
            if (index!== -1) {
                setIndex(index)
            } else {
                let editableNew = content.content.Editables
                editableNew.push({
                        "ID": id,
                        "type": 0,
                        "value": defaultText
                })
                content.setContent({
                    ...content.content,
                    Editable:editableNew
                })
                
                setIndex(content.content.Editables.length)
            }
        }
    },[content, content.content])
    if (!content.content || !content.content.Editables || index===-1) {
        return <Fragment>{defaultText}</Fragment>
    }
    return (content.administrate?<TextareaAutosize autoFocus style={{ boxSizing: "border-box", resize:"none" }} value={content.content.Editables[index]?content.content.Editables[index].value:defaultText} onChange={handleChange}  />:<Fragment>{parse(content.content.Editables[index]?content.content.Editables[index].value:defaultText)}</Fragment>)
}