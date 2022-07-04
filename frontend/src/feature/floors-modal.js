import { useContext, useEffect } from "react"
import { useState } from "react"
import { Link } from "react-router-dom"
import { RangeSlider } from "../components/range-slider"
import { ContentContext } from "../context/contentContext"
import { flatIDs } from "./flat-ids"

export const FloorsModal = ({title, classes, liter, opened, close}) => {
  
  const [entrance, setEntrance] = useState(0)
  const [floor, setFloor] = useState(1)
  const [flatAllIDs, setAllFlatIds] = useState([])
  const content = useContext(ContentContext)
  useEffect(()=>{
    if (content.content && content.content.Plans) {
      const ids = []
      content.content.Plans.map((flat)=>{
        ids.push(flat.ID)
      })
      setAllFlatIds(ids)
    }
  },[content.content])
  return (
    <div className={opened?`modal-full ${classes} open`:`modal-full ${classes}`}  >
      <div className="modal-full__close js-close-modal" onClick={()=>close()}>
        <div className="close-icon">
          <svg className="svg"><use xlinkHref="img/sprite.svg#close"></use></svg>
        </div>
      </div>
      <div className="modal-full__content">
        <div className="modal-flat__header">
          <div className="modal-flat-title">{title}</div>
          <div className="modal-flat-filter">
            <div className="floor-slider">
              <div className="floor-slider__title">Этаж</div>
              <RangeSlider value={floor} setValue={(floor) => setFloor(floor)} min={1} max={9} />
            </div>
            <ul className="entrance-list">
              <li className={entrance==0?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(0)}>подьезд 1</li>
              <li className={entrance==1?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(1)}>подьезд 2</li>
              <li className={entrance==2?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(2)}>подьезд 3</li>
            </ul>
          </div>
        </div>
        <div className={"modal-flat-map"}>
          <div id={"lit"+liter+"__flat-"+floor} className={floor>1?"lit"+(entrance+1)+"__flat-1__2-9":"lit"+(entrance+1)+"__flat-1__"+floor}>
            <img src={`img/entrances/${entrance+1}/${floor>1?"2-9":1}/floor.png`} alt="" />
            {flatIDs[`liter${liter}`] && flatIDs[`liter${liter}`][`floor${floor}`] && flatIDs[`liter${liter}`][`floor${floor}`][`entrance${entrance+1}`]?.map((id, index) => {
              if (id !=="0" && flatAllIDs.includes(id)){
                return  <Link to={"/plans/"+id}><div class={`flat-room flat-room_${index+1}`} ><div class="flat-room_before"></div></div></Link>
              }
            })}
          </div>
        </div>
        <div className="modal-flat-filter modal-flat-filter--mobile">
          <div className="floor-slider">
            <div className="floor-slider__title">Этаж</div>
            <RangeSlider value={floor} setValue={(floor) => setFloor(floor)} min={1} max={9} />
          </div>
          <ul className="entrance-list">
            <li lassName={entrance==0?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(0)}>подьезд 1</li>
            <li className={entrance==1?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(1)}>подьезд 2</li>
            <li className={entrance==2?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(2)}>подьезд 3</li>
          </ul>
        </div>
      </div>
    </div>
  )
}
