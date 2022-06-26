import { title } from "process";
import { FC } from "react";
import { Setting, Title } from "../../api";

interface TableItemProps {
  setting: Setting,
  index: number,
  editAction: any,
}

export const TableItem: FC<TableItemProps> = (props) => {
  return(<tr>
          <th scope="row">{props.index}</th>
          <td>{props.setting.name}</td>
          <td>{props.setting.value!.toString()}</td>
          <td className="col-sm" style={{width:"235px"}}>
            <div className="">
              <button className="btn btn-soft-success btn-sm" type="button" onClick={()=>{props.editAction()}}>Редактировать</button>
            </div>
          </td>
        </tr>
  )
}