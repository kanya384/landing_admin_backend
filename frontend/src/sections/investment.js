import { EditableText } from "../components/editable-text"

export const Investment = () => {
    return <div className="lvl-investing">
                <div className="wrapper">
                    <div className="about-info-item">
                        <div className="img"><img src="img/investing.png" alt="" /></div>
                        <div className="b-text">
                        <div className="title"><EditableText id={"62af012c2f3ab6f9b4a854de"} defaultText={"Выгодные инвестиции в ЖК AVAnta"}/></div>
                        <div className="text"><EditableText id={"62af012c2f3ab6f9b4a854df"} defaultText={"Увеличение спроса на недвижимость на Черноморском побережье, приводит к постоянному росту цен на рынке. Только за 2021 год новостройки Анапы подорожали на 21%. Помимо роста стоимости актива, покупатели квартир в Анапе могут получать пассивный доход от сдачи квартиры в аренду. "}/></div>
                        </div>
                    </div>
                    <div className="lvl-row-small">
                    <div className="lvl-row-small__img"><img src="img/investing2.png" alt="" /></div>
                    <div className="lvl-row-small__text"><EditableText id={"62af012c2f3ab6f9b4a854e0"} defaultText={"На сегодняшний день инвестиции  в недвижимость по-прежнему один из самых надежных способов сохранение и приумножения капитала."}/></div>
                </div>
                </div>
            </div>
}

export default Investment