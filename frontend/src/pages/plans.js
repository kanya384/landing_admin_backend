import React, { useContext, useEffect, useState } from "react"
import { Link, useLocation } from "react-router-dom";
import { AppBar } from "../components/appbar"
import { EditableText } from "../components/editable-text";
import { Footer } from "../components/footer"
import { RangeSlider } from "../components/range-slider";
import { ContentContext } from "../context/contentContext";

export const Plans = () => {
  const location = useLocation();
  const content = useContext(ContentContext)
  const [filtredPlans, setFiltredPlans] = useState()
  const [showParams, setShowParams] = useState()
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
  return( <React.Fragment>
      <AppBar />
      <div className="wrapper">
          <div className="page-title"><h1 className="h1"><EditableText id={"62aefdb322de07ce5e31751e"} defaultText={"выбрать квартиру"}/></h1></div>
          <div className="content-wrapper">
            <aside className="left-column">
              <div className="filter-left-btn-toggle">
                <div className="btn-toggle-filter js-btn-toggle-filter" onClick={()=>{setShowParams(!showParams)}}>{showParams?"Свернуть параметры":"Развернуть параметры"}</div>
              </div>
              <div className="filter-left js-filter-left-toggle" style={{display:showParams?"block":"none"}}>
                <div className="filter-item">
                  <div className="filter-item__title">Площадь, м²</div>
                  <div className="filter-item__body">
                    <div className="input-group-x2">
                      <div className="input-group-x2__col">
                        <div className="inp-group">
                          <div className="inp-group-label">От</div>
                          <input className="input" placeholder="От" name="areaMin" onChange={inputChange} value={filter.areaMin} />
                        </div>
                      </div>
                      <div className="input-group-x2__col">
                        <div className="inp-group">
                          <div className="inp-group-label">До</div>
                          <input className="input" placeholder="До" name="areaMax" onChange={inputChange} value={filter.areaMax} />
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="filter-item">
                  <div className="filter-item__title">Этаж</div>
                  <div className="filter-item__body">
                  <RangeSlider value={filter.floors} setValue={(floor) => setFloors(floor)} min={1} max={9} />
                  </div>
                </div>
                <div className="filter-item">
                  <div className="filter-item__body">
                    <ul className="checked-list">
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
                <div className="filter-btns">
                  <button className="btn-reset" onClick={resetFilter}>
                    <svg className="svg ico-reset"><use xlinkHref="img/sprite.svg#close"></use></svg> Сбросить
                  </button>
                </div>
              </div>
            </aside>
            <main className="content">
              <div className="listing-filter">
                <div className="listing-sort">
                  <div className="listing-sort__title">Сортировать</div>
                  <div className="listing-sort__select custom-select js-custom-select-wrapper">
                    <div className="custom-select__label js-custom-select-btn-toggle" onClick={()=>setShowDropDown(!showDropDown)}>{filter.sort==="area"?"по площади":filter.sort ==="price"?"по цене":"по этажу"}</div>
                    <div className="custom-select__arr js-custom-select-btn-toggle" onClick={()=>setShowDropDown(!showDropDown)}>
                      <svg className="svg ico-select-arr"><use xlinkHref="img/sprite.svg#select-arr"></use></svg>
                    </div>
                    <div className="custom-select__dropdown js-custom-select-dropdown" style={{display: showDropDown?"block":"none"}}>
                      <ul className="custom-select__list" style={{cursor:"pointer"}}>
                        <li className={filter.sort==="area"?"custom-select__item active":"custom-select__item"} onClick={()=>{setShowDropDown(false); setFilter({...filter, sort:"area"})}}>по площади</li>
                        <li className={filter.sort==="price"?"custom-select__item active":"custom-select__item"} onClick={()=>{setShowDropDown(false); setFilter({...filter, sort:"price"})}}>по цене</li>
                        <li className={filter.sort==="floor"?"custom-select__item active":"custom-select__item"} onClick={()=>{setShowDropDown(false); setFilter({...filter, sort:"floor"})}}>по этажу</li>
                      </ul>
                    </div>
                  </div>
                </div>
              </div>
              <div className="listing-choise">
                {filtredPlans&&filtredPlans.map((plan, index)=>{
                  if (index > 20) {
                    return
                  }
                  return <div className="choise-item">
                          <div className="choise-item__img"><Link to={"/plans/"+plan.ID}><img src={process.env.REACT_APP_BACKEND+"store/"+plan.image} alt="" /></Link></div>
                          <div className="choise-item__right">
                            <div className="choise-item__header">
                              <div className="choise-item__title"><a href="#">{plan.rooms}-комн, {plan.area} м²</a></div>
                              <div className="choise-item__floor">{plan.floor} этаж из 9</div>
                              <div className="choise-item__floor">{plan.entrance} подьезд</div>
                              <div className="choise-item__price">1 000 000 руб.</div>
                            </div>
                            <div className="choise-item__bottom">
                              <div className="choise-item__link-pdf"><a href="#" className="link-pdf">Скачать в PDF</a></div>
                              <div className="choise-item__more"><Link className="link-bb" to={"/plans/"+plan.ID}>узнать стоимость</Link></div>
                            </div>
                          </div>
                        </div>
                })}
              </div>
            </main>
          </div>
        </div>
        <Footer />
    </React.Fragment>
  )
}

export default Plans
