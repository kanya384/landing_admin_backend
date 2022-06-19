import React, { useContext, useEffect, useState } from "react"
import { useSearchParams } from "react-router-dom";
import { ContentContext } from "../context/contentContext";
import Cookies from 'js-cookie';
import "./panel.css"
import axios from 'axios';

export const Panel = () => {
    const content = useContext(ContentContext)
    const [showPanel, setShowPanel] = useState(false)

    const [searchParams, setSearchParams] = useSearchParams();
    useEffect(()=>{
        if (searchParams.get("administarate") === "y") {
            checkAuth(()=>{content.setAdministrate(true)})
        }
    },[searchParams])

    const checkAuth = (callback) => {
        let token = getTokenFromCookies()
        fetch(process.env.REACT_APP_BACKEND+"ping", {headers: {"Authorization": token}})
            .then((response) => {
                if (response.status === 200) {   
                    callback()
                }
            })
    }

    const storeEditables = () => {
        let token = getTokenFromCookies()
        let data = []
        content.content.Editables.forEach((editable)=>{
            data.push({
                "id": editable.ID,
                "type": 0,
                "value": editable.value,
            })
        })
        let sendData = JSON.stringify(data)
        axios.post(process.env.REACT_APP_BACKEND+"editable", data, {
            headers: {
                'Content-Type':'application/json',
                "Authorization": token,
                'Accept':'application/json'
            }
        })
    }

    const exit = () => {
        console.log('exit')
        window.location.href = "http://"+window.location.host
    }

    const getTokenFromCookies = () => {
        let token = ""
        let tmp = Cookies.get("token")
        if (tmp) {
          token = tmp
        }
        return token
    }
    
   
    return ( 
                <React.Fragment>
                    <section className="edit-panel" style={{display:content.administrate?"flex":"none"}}>
                        <span className="edit-panel-text">Режим редактирования</span>
                        <div className="edit-panel-buttons">
                            <button onClick={exit}>Выйти</button>
                            <button onClick={storeEditables}>Сохранить</button>
                        </div>
                    </section>
                    <div className="space-bottom" style={{marginTop:"30px"}}></div>
                </React.Fragment>
    )
}