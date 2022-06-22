import { useContext, useState } from "react"
import { Link } from "react-router-dom"
import { SectionsContext } from "../context/sectionsContext"

export const Menu = () => {
  const blocks = useContext(SectionsContext)
  const [showMobile, setShowMobile] = useState(false)

  const clickMenu = (e) => {
      e.preventDefault()
      let popup = e.currentTarget.getAttribute("href")
      if (blocks.blocks < 19) {
          blocks.setBlocks(19)
          setTimeout(() => {
            window.scrollTo({
              top: document.querySelector("." + popup).offsetTop,
              behavior: "smooth"
            })
          }, 1200)

      } else {
          window.scrollTo({
              top: document.querySelector("." + popup).offsetTop,
              behavior: "smooth"
          })
      }
      setShowMobile(false)
  }

  const hamburger = (e) =>{
    e.preventDefault()
    if (!document.querySelector('body').classList.contains('active_m')){
      document.querySelector('body').classList.add('active_m')
    } else {
      document.querySelector('body').classList.remove('active_m')
    }
    setShowMobile(!showMobile)
  } 


  return(
    <header className="header">
      <div className="wrapper">
        <div className="logo">
          <a href="#">
            <img src="./img/logo.svg" alt="logo" />
          </a>
        </div>
        <div className="header-phone">
          <div className="phone">+7 987 654-32-10</div>
          <a href="#" className="btn-recall">Заказать звонок</a>
        </div>
        
        <a href="#" id="hamburger-icon" onClick={hamburger}>
          <span className="line line-1"></span>
          <span className="line line-2"></span>
          <span className="line line-3"></span>
          <strong>МЕНЮ</strong>
        </a>
      <nav className="navigation" style={{display:showMobile?"block":"none"}}>
        <ul>
          <li><a onClick={clickMenu} href="lvl2">Инфраструктура</a></li>
          <li><a onClick={clickMenu} href="lvl8">Планировки и цены</a></li>
          <li><a onClick={clickMenu} href="lvl9">Способы покупки</a></li>
          <li><Link to="/about">О застройщике</Link></li>
          <li><a onClick={clickMenu} href="footer">Контакты</a></li>
        </ul>
      </nav>
    </div>
  </header>
  )
}