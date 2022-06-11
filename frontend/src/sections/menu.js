import { Link } from "react-router-dom"

export const Menu = () => {
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
        
        <a href="#" id="hamburger-icon">
          <span className="line line-1"></span>
          <span className="line line-2"></span>
          <span className="line line-3"></span>
          <strong>МЕНЮ</strong>
        </a>
      <nav className="navigation">
        <ul>
          <li><a href="">Инфраструктура</a></li>
          <li><a href="">Планировки и цены</a></li>
          <li><a href="">Способы покупки</a></li>
          <li><Link to="/about">О застройщике</Link></li>
          <li><a href="">Контакты</a></li>
        </ul>
      </nav>
    </div>
  </header>
  )
}