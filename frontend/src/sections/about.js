import React, { useContext, useState } from "react"
import { EditableText } from "../components/editable-text";
import { Modal } from "../components/modals"
import { ContentContext } from "../context/contentContext";

export const  About = () => {
  const content = useContext(ContentContext)
  const [isOpen, setModalState] = useState(null)
  const [info, setInfo] = useState({
    title: "",
    description:"",
  })
  return ( <React.Fragment>
          <div className="lvl6">
            <div className="wrapper">
              <div className="about-title"><EditableText id={"62aef61ba26e626025a8d8ca"} defaultText={"О проекте"}/></div>
              <div className="about-text">
                <EditableText id={"62aef61ba26e626025a8d8c9"} defaultText={"3 девятиэтажных дома на Высоком берегу по своему архитектурному облику напоминают волны, среди которых расположится благоустройство и озеленение жилого комплекса, в традиционно высоком исполнении"}/>
              </div>
              <div className="about-info-list">
                {content.content && content.content.ProjectInfo && content.content.ProjectInfo.map((info) => {
                  return  <div className="about-info-item">
                              <div className="img"><img src={process.env.REACT_APP_BACKEND+"store/"+ info.photo} alt="" /></div>
                              <div className="b-text">
                                <div className="title">{info.title}</div>
                                <div className="text">{info.anonce}</div>
                                <div className="podr"><a href="" onClick={(e) => {e.preventDefault();  setInfo(info); setModalState(true)}}>Подробнее</a></div>
                              </div>
                          </div>
                })}
              </div>
            </div>
          </div>
          <Modal 
            title={info.title}
            content={info.description}
            position={window.pageYOffset}
            fields={[]}
            celtype={""}
            image ={"img/news-img.svg"}
            opened={isOpen}
            close = {()=>{setModalState(null)}}
        />
    </React.Fragment>
  )
}

export default About