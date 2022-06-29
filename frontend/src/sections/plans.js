import React, { useEffect, useState } from "react"
import { Modal } from "../components/modals"
import { Link } from "react-router-dom";
import Konva from "konva"
import { EditableText } from "../components/editable-text";
var stage
var layer


export const Plans = () => {
  const [isOpen, setModalState] = useState(false)
  const [hovered, setHovered] = useState(false)
  const [selectedLiter, setLiter] = useState()
  const paths = [
		{
			name: 'paths_plan1',
			liter: 3,
			path: 'M1390.09 339.381L1343.58 476.001V483.327H1334.72L1326.38 490.257H1272.24L1268.56 482.028L1259.03 483.327L1248.2 491.773H1226.11V480.945L1132.7 485.767L1125.75 467.016H1115.18L1108.23 476.058L1103.67 476.725L1078.4 498.465L1077.39 503.106L1075.07 506.731L1081.88 512.097L1008.1 578.445L1005.64 577.44L1000.72 581.798L989.474 572.728L918.39 631.474H855.945L825.367 644.831H811.638L806.972 642.102H759.658L756.726 644.831H740.047L736.913 642.102H727.706V634.852H720.835L716.111 639.79L708.811 643.655H692.064V618.749L687.038 617.89V611.496L682.759 588.339H644.248L617.777 581.543V553.014L610.154 446.764L614.305 410.772H642.205V404.717H694.845V415.549H716.154V410.791H720.355L721.773 408.412H733.92L735.866 410.511H828.106L832.458 407.679L840.318 407.933L843.361 409.784H851.162V421.532H891.273L951.553 379.538L939.038 370.375L996.821 330.682L1009.67 338.358L1075.93 289.583L1081.18 294.611L1092.69 284.852H1161.14V279.612L1166.28 276.678H1174.38L1178.26 279.612H1225.53L1229.49 276.678H1236.02L1240.72 277.852L1246.73 282.766H1268.66V279.612L1272.11 276.678H1277.68L1284.14 277.852L1290.09 283.94H1319.66L1327.93 298.313L1347.75 297.322L1355.68 310.373H1371.38L1390.09 339.381Z',
		},
		{
			name: 'paths_plan2',
			liter: 1,
			path: 'M727.672 167.187V43.8598L714.699 43.4685V35.0264H700.931V27.7254H678.722V24.1232H667.423L661.708 26.9489H647.557L641.718 24.1232H633.956H594.986L588.419 23.1471L579.989 27.7254H528.501L515.957 33.6284L513.375 31.4148L452.685 56.1336L448.995 53.551L442.539 56.1336L436.82 53.551L393.655 71.0756H389.412L378.713 78.4543L381.295 82.6971L329.644 102.62H300.867L302.712 97.2701H265.634L261.391 100.406L258.255 93.0273H233.536L230.769 97.2701H216.38H205.866L200.516 100.406H186.128L187.972 94.6875H151.632L149.787 98.3769L128.204 97.2701L113.816 117.746L150.92 238.162L148.289 245.304L156.934 247.935V256.204H165.203L167.083 265.226H179.111L188.884 269.401L252.442 267.701L260 262H319.5L333.5 229H356L357.5 224.5H403L397 232H415L417 227H438L481.709 198.21L498.686 208.578L509.619 207.052L518.263 200.442L513.941 194.34L547.247 178.323H568.348V182.472L583.603 186.204L601.41 174.308H635.959L649.507 184.036L680.426 182.472L727.672 167.187Z',
		},
		{
			name: 'paths_plan3',
			liter: 2,
			path: 'M349.147 383.5L322.647 254.5L333.147 229H356.147L357.647 224.5H402.647L397.147 232H415.147L417.147 227H433.147L436.147 229H455.147L458.147 224H485.647L486.647 232L488.147 229H532.647V235.5H564.647L615.647 209.5V199L675.647 170.5L683.647 175.5L753.647 139L759.647 142L768.647 136H827.147L830.147 131L897.147 129.5L905.647 133.5L918.647 133L923.147 131L931.647 129L938.647 131L942.647 134.5H966.147L969.647 144.5H984.647L988.147 155H1001.65L1004.65 161L1018.15 168L1009.65 177.5L992.647 310.5H974.647L969.647 313L899.147 310.5L885.147 318H849.147L837.647 321L811.147 318L807.147 310.5H790.647L674.647 378.5L665.647 375.5L596.647 403.5H551.647L539.147 417H527.147L524.147 426L509.147 429L502.147 426H455.647H441.147L435.647 417L416.147 426H396.147V417L372.647 414.5L349.147 383.5Z',
		},
	]

  useEffect(() => {
		var stageWidth = 1500;
		var stageHeight = 704;
		stage = new Konva.Stage({
			container: 'paths_plan',
			width: stageWidth,
			height: stageHeight,
		});
		layer = new Konva.Layer();
		stage.add(layer)
		paths.map((el) => {
			var path = new Konva.Path({
				opacity: 0.7,
				data: el.path,
				fill: hovered == el.liter ? el.condition == 2 ? 'yellow' : '#0CA5E6' : 'transparent',
				scaleX: 1,
				scaleY: 1
			});
			path.on('mouseenter', function (e) {
				var fill = el.condition == 2 ? 'green' : '#0CA5E6';
				this.fill(fill);
				layer.draw();
				setHovered(el.liter)
				const container = e.target.getStage().container();
				container.style.cursor = "pointer"
			})
			path.on('mouseleave', function (e) {
				var fill = 'transparent';
				this.fill(fill);
				layer.draw();
				setHovered(false)
				const container = e.target.getStage().container();
				container.style.cursor = "default"
			})
			path.on('click', function () {
				const liter = el.liter
        setLiter(liter)
        setModalState(true)
			})
			path.on('touchstart', function () {
				const liter = el.liter
        setLiter(liter)
        setModalState(true)
				
			})
			layer.add(path)
		})
		layer.draw()

		function fitStageIntoParentContainer() {
			var container = document.querySelector('#paths_plan');
			var containerWidth = container?.offsetWidth;
			var scale = containerWidth / stageWidth;
			stage.width(stageWidth * scale);
			stage.height(stageHeight * scale);
			stage.scale({ x: scale, y: scale });
			stage.draw();
		}

		fitStageIntoParentContainer();
    setTimeout(()=>{
      fitStageIntoParentContainer();
    },500)
		// adapt the stage on any window resize
		window.addEventListener('resize', fitStageIntoParentContainer);
	})



  return (
    <React.Fragment>
      <div className="lvl8">
        <div className="wrapper">
          <div className="h-title"><EditableText id={"62af012c2f3ab6f9b4a854da"} defaultText={"Подбор квартиры на 3D-плане"}/></div>
          <div className="text"><EditableText id={"62af012c2f3ab6f9b4a854db"} defaultText={"Выберите дом и этаж, чтобы посмотреть планировки и узнать цены. Передвигайтесь влево-вправо, ввех и вниз"}/></div>
          <div className="b-3d">
			  <div className="b-3d__container">
				  <img src="img/home3d.jpg" alt="" />
				  <div className="plan-float-label plan-float-label__1" onClick={()=>{setLiter(1); setModalState(true)}}>
					  <div className="liter">
						  литер 1
						  <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
							  <path d="M12.501 4.5L18.501 10.5M18.501 10.5L12.501 16.5M18.501 10.5L1.50098 10.5" stroke="white" stroke-width="2"/>
						  </svg>
					  </div>
					  <div className="date">3 кв. 2023</div>
				  </div>
				  <div className="plan-float-label plan-float-label__2" onClick={()=>{setLiter(1); setModalState(true)}}>
					  <div className="liter">
						  литер 2
						  <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
							  <path d="M12.501 4.5L18.501 10.5M18.501 10.5L12.501 16.5M18.501 10.5L1.50098 10.5" stroke="white" stroke-width="2"/>
						  </svg>
					  </div>
					  <div className="date">3 кв. 2023</div>
				  </div>
				  <div className="plan-float-label plan-float-label__3" onClick={()=>{setLiter(1); setModalState(true)}}>
					  <div className="liter">
						  литер 3
						  <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
							  <path d="M12.501 4.5L18.501 10.5M18.501 10.5L12.501 16.5M18.501 10.5L1.50098 10.5" stroke="white" stroke-width="2"/>
						  </svg>
					  </div>
					  <div className="date">3 кв. 2023</div>
				  </div>
				  <div id="paths_plan"></div>
			  </div>
          </div>
          
          <div className="b-link-row">
            <Link to="/plans" className="lnk-params js-open-modal-flat">
              <span className="lnk-params__ico">
              <svg className="ico" width="66" height="30" viewBox="0 0 66 30"
                    xmlns="http://www.w3.org/2000/svg">
                <path d="M64.5 14.9983L65.2033 15.7092L65.9219 14.9982L65.2032 14.2874L64.5 14.9983ZM51.0509 29.7109L65.2033 15.7092L63.7967 14.2875L49.6443 28.2891L51.0509 29.7109ZM65.2032 14.2874L51.0508 0.289034L49.6444 1.71097L63.7968 15.7093L65.2032 14.2874ZM64.5 13.9983H0.5V15.9983H64.5V13.9983Z" />
              </svg>
              </span>
			  <EditableText id={"62af012c2f3ab6f9b4a854e6"} defaultText={"Подбор по параметрам"}/>
            </Link>
          </div>
        </div>
      </div>
      <Modal
        title={`Планировки литера ${selectedLiter}`}
        position={window.pageYOffset}
        classes={"modal-flat"}
        liter={selectedLiter}
        fields={[
            {
                type:"text",
                name: "name",
                placeholder: "Имя",
                required: false,
            },
            {
                type:"text",
                name: "phone",
                placeholder: "Телефон",
                required: true,
            },
        ]}
        btnTitle={"Отправить"}
        celtype={"getPresentation"}
        image ={"img/get-present-bg.svg"}
        imageMobile={"img/get-present-bg-mobile.svg"}
        opened={isOpen}
        close = {()=>{setModalState(null)}}
      />
    </React.Fragment>
  )
}

export default Plans
