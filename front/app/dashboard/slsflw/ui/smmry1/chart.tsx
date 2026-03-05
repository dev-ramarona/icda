"use client";
import { Area, AreaChart, Bar, CartesianGrid, ComposedChart, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts";
export default function UixSlsflwSmmry1Charts({ arrdta }: { arrdta: { name: string; sales: number }[] }) {
    return (
        <div className="w-full h-52">
            <ComposedChart
                // style={{ width: '100%', maxWidth: '700px', maxHeight: '70vh', aspectRatio: 1.618 }}
                className="w-full h-full"
                responsive
                data={arrdta}
                margin={{
                    right: 20,
                    left: 20,
                    top: 10
                }}
            >
                <CartesianGrid stroke="#f5f5f5" />
                <XAxis dataKey="name" />
                <YAxis width="auto" />
                <Tooltip />
                <Legend />
                <Bar dataKey="sales" barSize={20} fill="#00598a" />
                <Line type="monotone" dataKey="sales" stroke="#0092b8" />
            </ComposedChart>
            {/* <ResponsiveContainer>
                <LineChart data={arrdta} margin={{ right: 20, left: 20, }}>
                    <CartesianGrid strokeDasharray="3 10" />
                    <XAxis dataKey="name" />
                    <Tooltip />
                    <Line type="monotone" dataKey="sales" stroke="#8884d8" />
                </LineChart>
            </ResponsiveContainer> */}
        </div>
    );

}
