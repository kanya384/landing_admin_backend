import { Month } from "../../api";

interface MonthSelect {
    value: number;
    name: string;
    type: number;
}
export const months: MonthSelect[] = [
    {
        value: 0,
        name: "Январь",
        type: 0,
    },
    {
        value: 1,
        name: "Февраль",
        type: 0,
    },
    {
        value: 2,
        name: "Март",
        type: 0,
    },
    {
        value: 3,
        name: "Апрель",
        type: 0,
    },
    {
        value: 4,
        name: "Май",
        type: 0,
    },
    {
        value: 5,
        name: "Июнь",
        type: 0,
    },
    {
        value: 6,
        name: "Июль",
        type: 0,
    },
    {
        value: 7,
        name: "Август",
        type: 0,
    },
    {
        value: 8,
        name: "Сентябрь",
        type: 0,
    },
    {
        value: 9,
        name: "Октябрь",
        type: 0,
    },
    {
        value: 10,
        name: "Ноябрь",
        type: 0,
    },
    {
        value: 11,
        name: "Декабрь",
        type: 0,
    },
]

export const excludeExistingMonths = (monthsList:MonthSelect[], existingMonths: Month[]):MonthSelect[] => {
    let resultList:MonthSelect[] =[]
    monthsList.map((month) => {
        let flag = true
        for (let i=0; i< existingMonths.length; i++) {
            if (month.value === existingMonths[i].value || (existingMonths[i].value === undefined && month.value=== 0)){
                flag= false
                break;
            }
        }
        if (flag) {
            resultList.push(month)
        }
    })
    return resultList
}