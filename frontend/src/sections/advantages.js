import { SliderWithText } from "../components/slider-with-text";

export const Advantages = () => {
  let slides1 = {
    title: "Морские развлечения",
    text: "Море - самое место для отдыха! Жителям Анапы каждый день и каждые выходные доступны морские прогулки, поездки на бананах и таблетках, дайвинг и морская рыбалка. Присоединяйтесь к ним!",
    slides: [
      "img/img10.png",
      "img/img11.png",
      "img/img12.png",
    ]
  }
  let slides2 = {
    title: "Воздух с ароматом можжевельника",
    text: "Мощный оздоровительный эффект дает ионизированный воздух многовековых реликтовых можжевеловых лесов, расположенных в районе заповедника Большой Утриш и села Сукко",
    slides: [
      "img/img11.png",
      "img/img12.png",
      "img/img10.png",
    ]
  }

  let slides3 = {
    title: "Бесконечный песчаный пляж",
    text: "Линия песчаного берега растянулась на 42 км. Всё это барханы и дюны с целебным кварцевым песком, подобных которым не найти ни в нашей стране, ни в Европе",
    slides: [
      "img/img12.png",
      "img/img10.png",
      "img/img11.png",
    ]
  }

  return( <div className="lvl11">
            <div className="wrapper">
              <h2 className="lvl-title">А также вас ждёт:</h2>
              <div className="about-info-list">
                <SliderWithText slides={slides1} key={"slider1"} />
                <SliderWithText slides={slides2} key={"slider2"}/>
                <SliderWithText slides={slides3} key={"slider3"}/>
              </div>  
            </div>
          </div>
  )
}