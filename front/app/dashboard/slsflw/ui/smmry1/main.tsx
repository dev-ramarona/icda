import UixSlsflwSmmry1Charts from "./chart";

export default async function UixSlsflwSmmry1Mainpg() {
    const arrdta = [
        { name: "19-Jan", sales: 500 },
        { name: "18-Jan", sales: 600 },
        { name: "17-Jan", sales: 400 },
        { name: "16-Jan", sales: 300 },
        { name: "15-Jan", sales: 200 },
        { name: "14-Jan", sales: 700 },
        { name: "13-Jan", sales: 600 },
        { name: "12-Jan", sales: 300 },
        { name: "11-Jan", sales: 200 },
        { name: "10-Jan", sales: 400 },
        { name: "09-Jan", sales: 400 },
        { name: "08-Jan", sales: 600 },
        { name: "07-Jan", sales: 500 },
        { name: "06-Jan", sales: 500 },
        { name: "05-Jan", sales: 400 },
        { name: "04-Jan", sales: 400 },
        { name: "03-Jan", sales: 100 },
        { name: "02-Jan", sales: 600 },
        { name: "01-Jan", sales: 600 },
    ];
    return (
        <>
            {arrdta.length > 0 ? (
                <UixSlsflwSmmry1Charts arrdta={arrdta} />
            ) : (
                <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
                    No database Log Action
                </div>
            )}
        </>
    );
}