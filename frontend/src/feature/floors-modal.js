import { useContext, useEffect } from "react"
import { useState } from "react"
import { Link, useSearchParams } from "react-router-dom"
import { RangeSlider } from "../components/range-slider"
import { ContentContext } from "../context/contentContext"
import { SectionsContext } from "../context/sectionsContext"
import { flatIDs } from "./flat-ids"

export const FloorsModal = ({title, classes, liter, opened, close}) => {
  
  const [entrance, setEntrance] = useState(0)
  const [floor, setFloor] = useState(1)
  const [flatAllIDs, setAllFlatIds] = useState([])
  const content = useContext(ContentContext)
  const sections = useContext(SectionsContext)
  const [searchParams, setSearchParams] = useSearchParams();
  useEffect(()=>{
    if (content.content && content.content.Plans) {
      const ids = []
      content.content.Plans.map((flat)=>{
        ids.push(flat.ID)
      })
      setAllFlatIds(ids)
    }
  },[content.content])

  useEffect(()=>{
    setTimeout(()=>{
        if (searchParams.get("entrance")) {
				  setEntrance(parseInt(searchParams.get("entrance")))
        }
        if (searchParams.get("floor")) {
				  setFloor(parseInt(searchParams.get("floor")))
        }
        if (searchParams.get("floor")) {
          try{
              searchParams.delete("entrance")
              searchParams.delete("floor")
              searchParams.delete("liter")
              setSearchParams(searchParams)
              setTimeout(()=>{
                window.scrollTo({
                  top: document.querySelector(".lvl8").offsetTop,
                })
              },[500])
             
          } catch(e){
            console.log(e)
          }
        }
      }, [500])
    },[])

  const flatClick = (id)=>{
    sections.setShowPlan(id)
  }
    
    
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
              <div className="floor-slider__title">????????</div>
              <RangeSlider value={floor} setValue={(floor) => setFloor(floor)} min={1} max={9} />
            </div>
            <ul className="entrance-list">
              <li className={entrance==0?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(0)}>?????????????? 1</li>
              <li className={entrance==1?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(1)}>?????????????? 2</li>
              <li className={entrance==2?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(2)}>?????????????? 3</li>
            </ul>
          </div>
        </div>
        <div className={"modal-flat-map"}>
          <div id={"lit"+liter+"__flat-"+floor} className={floor>1?"lit"+(entrance+1)+"__flat-1__2-9":"lit"+(entrance+1)+"__flat-1__"+floor}>
            <img src={`img/entrances/${entrance+1}/${floor>1?"2-9":1}/floor.png`} alt="" />
            {flatIDs[`liter${liter}`] && flatIDs[`liter${liter}`][`floor${floor}`] && flatIDs[`liter${liter}`][`floor${floor}`][`entrance${entrance+1}`]?.map((id, index) => {
              if (id !=="0" && flatAllIDs.includes(id)){
                return  <div class={`flat-room flat-room_${index+1}`} onClick={()=>flatClick(id)}><div class="flat-room_before"></div></div>// <Link to={"/plans/"+id+window.location.search}></Link>
              }
            })}
          </div>
        </div>
        <div className="modal-flat-filter modal-flat-filter--mobile">
          <div className="floor-slider">
            <div className="floor-slider__title">????????</div>
            <RangeSlider value={floor} setValue={(floor) => setFloor(floor)} min={1} max={9} />
          </div>
          <ul className="entrance-list">
            <li className={entrance==0?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(0)}>?????????????? 1</li>
            <li className={entrance==1?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(1)}>?????????????? 2</li>
            <li className={entrance==2?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(2)}>?????????????? 3</li>
          </ul>
        </div>
      </div>
    </div>
  )
}
