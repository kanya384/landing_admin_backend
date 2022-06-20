import { EditableText } from "../components/editable-text"
import { RangeSlider } from "../components/range-slider";
import React, {useContext, useEffect, useState} from "react";
import {useLocation} from "react-router-dom";
import {ContentContext} from "../context/contentContext";

export const GetPodbor = () => {
  const location = useLocation();
  const content = useContext(ContentContext)
  const [filtredPlans, setFiltredPlans] = useState()
  const [filter, setFilter] = useState({
    areaMin: "",
    areaMax: "",
    floors: [1,9],
    rooms: [1,2,3],
    liters: [1,2,3],
    sort: "area",
  })
  const [showDropDown, setShowDropDown] = useState(false)
  
  const filterPlans = () => {
    let min = filter.areaMin!==""?filter.areaMin:0
    let max = filter.areaMax!==""?filter.areaMax:999
    let plans = []
    content.content?.Plans.map((plan)=> {
      if (plan.area>=min && plan.area<=max && plan.floor >= filter.floors[0] && plan.floor <= filter.floors[1] && filter.rooms.includes(plan.rooms) && filter.liters.includes(parseInt(plan.liter.replace('ЖК ВЫСОКИЙ БЕРЕГ Литер ',"")))){
        plans.push(plan)
      }
    })
    
    setFiltredPlans(sortFlats(filter.sort, plans))
  }
  
  const inputChange = (event) => {
    setFilter({
      ...filter,
      [event.target.name]: event.target.value.replace(",",".")
    })
  }
  
  const setFloors = (value) => {
    setFilter({
      ...filter,
      floors: value
    })
  }
  
  const setRooms = (event) => {
    let rooms = parseInt(event.currentTarget.getAttribute("rooms"))
    if (event.target.checked && !filter.rooms.includes(rooms)) {
      setFilter({
        ...filter,
        rooms: [
          ...filter.rooms,
          rooms
        ]
      })
    } else {
      setFilter({
        ...filter,
        rooms: filter.rooms.filter(room => room!=rooms)
      })
    }
  }
  
  const setLiters = (event) => {
    let liter = parseInt(event.currentTarget.getAttribute("liter"))
    if (event.target.checked && !filter.liters.includes(liter)) {
      setFilter({
        ...filter,
        liters: [
          ...filter.liters,
          liter
        ]
      })
    } else {
      setFilter({
        ...filter,
        liters: filter.liters.filter(liters => liter!=liters)
      })
    }
  }
  
  const resetFilter = (e) => {
    setFilter({
      areaMin: "",
      areaMax: "",
      floors: [1,9],
      rooms: [1,2,3],
      liters: [1,2,3],
    })
  }
  
  const sortFlats = (filedName, flats) => {
    
    return flats.sort((a,b)=>{
      if (a[filedName] > b[filedName]) {
        return 1;
      }
      if (a[filedName] < b[filedName]) {
        return -1;
      }
      return 0;
    })
  }
  
  useEffect(() => {
    window.scrollTo(0, 0);
  }, [location]);
  
  useEffect(()=>{
    if (content.content?.Plans){
      setFiltredPlans(content.content?.Plans)
    }
  },[content.content?.Plans])
  
  useEffect(()=>{
    if (filter){
      filterPlans()
    }
  },[filter])
  return(
    <div className="lvl12">
      <div className="wrapper">
        <div className="podborka-left">
          <div className="title"><EditableText id={"62aef61ba26e626025a8d8da"} defaultText={"Получите подборку свободных квартир с планировками и ценами"}/></div>
          <div className="bottom-text">
            <EditableText id={"62aef61ba26e626025a8d8db"} defaultText={"Перезвоним и предоставим самые выгодные предложения на сегодня"}/>
          </div>
        </div>
        <div className="podborka-right">
          <div className="calculator-section">
            <div className="calculator-filter">
              <div className="filter-item filter-item__floor">
                <div className="filter-item__title">Этаж</div>
                <div className="filter-item__body">
                  <RangeSlider value={filter.floors} setValue={(floor) => setFloors(floor)} min={1} max={9} />
                </div>
              </div>
              <div className="filter-item filter-item__rooms">
                <div className="filter-item__body">
                  <ul className="checked-list checked-list--column-3">
                    <li>
                      <label className="inp-checkbox"><input id="ck1" type="checkbox" rooms={1} onChange={setRooms} checked={filter.rooms.includes(1)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">1-комн.</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input id="ck2" type="checkbox" liter={1} onChange={setLiters} checked={filter.liters.includes(1)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">Корпус 1</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input type="checkbox" rooms={2} onChange={setRooms} checked={filter.rooms.includes(2)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">2-комн.</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input type="checkbox" liter={2} onChange={setLiters} checked={filter.liters.includes(2)}/>
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">Корпус 2</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input type="checkbox" rooms={3} onChange={setRooms} checked={filter.rooms.includes(3)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">3-комн.</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input type="checkbox" liter={3} onChange={setLiters} checked={filter.liters.includes(3)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">Корпус 3</span></label>
                    </li>
                  </ul>
                </div>
              </div>
              <div className="filter-item">
                <div className="form-ec form-ec--calculator">
                  <div className="form-ec__content">
                    <div className="form-ec__input-row">
                      <div className="inp-group">
                        <div className="inp-group-label">Ваше имя</div>
                        <input className="input" placeholder="Ваше имя" type="text" />
                      </div>
                      <div className="inp-group">
                        <div className="inp-group-label">Номер телефона</div>
                        <input className="input" placeholder="Номер телефона" type="text" />
                      </div>
                      <div className="inp-group">
                        <div className="inp-group-label">E-mail</div>
                        <input className="input" placeholder="E-mail" type="text" />
                      </div>
                      <div className="inp-group">
                        <button className="btn-submit"><span className="btn-submit-text">Отправить</span></button>
                      </div>
                    </div>
                    
                  </div>
                </div>
              </div>
            </div>
          </div>
          
        </div>
      </div>
    </div>
  )
}
