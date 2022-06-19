import React from "react"
import { EditableText } from "../components/editable-text"

export const Walks = () => {
  return ( 
        <React.Fragment>
            <div className="lvl3">
              <div className="wrapper">
                <div className="lvl-row">
                  <div className="lvl-col b-small">
                    <div className="title"><EditableText id={"62aef61ba26e626025a8d8be"} defaultText={"Прогулки на высоком берегу"}/></div>
                    <div className="text">
                      <EditableText id={"62aef61ba26e626025a8d8bd"} defaultText={"В Анапе очень много скверов, парков и других интересных и симпатичных мест. И множество из них находятся на Высоком берегу"}/>
                    </div>
                  </div>
                  <div className="lvl-col b-big">
                    <div className="img-float-label">
                      <img src="img/img3.jpg" alt="" />
                      <div className="b-label"><EditableText id={"62aef61ba26e626025a8d8c5"} defaultText={"Ореховая роща"}/></div>
                    </div>
                    <div className="b-text">
                      <EditableText id={"62aef61ba26e626025a8d8bf"} defaultText={"Парк с самым большим на Черноморском побережье розарием, включающим более 6 тысяч кусов роз разных сортов и расцветок"}/>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className="lvl4">
              <div className="wrapper">
                <div className="b-small-fix-height">
                  <div className="img-float-label">
                    <img src="img/img4.jpg" alt="" />
                    <div className="b-label"><EditableText id={"62aef61ba26e626025a8d8c1"} defaultText={"Анапский парк"}/></div>
                  </div>
                  <div className="text">
                    <EditableText id={"62aef61ba26e626025a8d8c2"} defaultText={"Достопримечательность и отличное место для прогулок и фото. Отсюда открывается восхитительный вид на бесконечную синеву моря с белыми кораблями и парусниками. Рядом с маяком много лавочек, беседо и кафе"}/>
                  </div>
                </div>
                <div className="b-small-fix-height">
                  <div className="img-float-label">
                    <img src="img/img5.jpg" alt="" />
                    <div className="b-label"><EditableText id={"62aef61ba26e626025a8d8c3"} defaultText={"Бювет минеральных вод"}/></div>
                  </div>
                  <div className="text">
                    <EditableText id={"62aef61ba26e626025a8d8c4"} defaultText={"Еще одно знаковое городское место. Это облагороженный источник, где можно попробовать знаменитую лечебно-стоповую минеральную воду «Анапская»"}/>
                  </div>
                </div>
              </div>
            </div>
        </React.Fragment>
  )
}