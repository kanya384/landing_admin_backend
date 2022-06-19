import { LightgalleryProvider, LightgalleryItem, useLightgallery} from "react-lightgallery";
import "lightgallery.js/dist/css/lightgallery.css";
import { ContentContext } from "../context/contentContext";
import { useContext } from "react";
import { EditableText } from "../components/editable-text";

export const Video = () => {
  const content = useContext(ContentContext)
  return ( <div className="lvl10 parrallax-scene2-end">
              <div className="wrapper">
                <div className="b-video">
                  <div className="b-text"><EditableText id={"62aef61ba26e626025a8d8d9"} defaultText={"Да, это Краснодарский край и его жемчужина - Анапа. Этот город ассоциируется с релаксом и праздником. Даже городская администрация находится рядом с пляжем"}/></div>
                  {content.content?<LightgalleryProvider>
                    <LightgalleryItem group={"video"} src={content.content.Video[0].url}>
                    <a href="#" className="b-video-player" onClick={(e)=>{e.preventDefault()}}>
                      <span className="play">
                        <svg width="30" height="40" viewBox="0 0 30 40" xmlns="http://www.w3.org/2000/svg">
                          <path d="M30 20L1.73794e-07 40L1.92225e-06 -1.31134e-06L30 20Z" />
                        </svg>
                      </span>
                      <span className="img">
                        <img src="img/circ-anapa.svg" alt="" />
                      </span>
                      <span className="label">
                        <span>Смотреть видео</span>
                      </span>
                    </a>
                    </LightgalleryItem>
                  </LightgalleryProvider>:""}
                </div>
              </div>
            </div>
  )
}