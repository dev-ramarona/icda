export default function UixGlobalTheadxTablex({
  firsth,
  mainhd,
}: {
  firsth: string;
  mainhd: string[];
}) {
  return (
    <thead>
      <tr>
        {firsth ? <th className="sticky left-0">{firsth}</th> : ""}
        {mainhd.map((key) => (
          <th key={key}>{key}</th>
        ))}
      </tr>
    </thead>
  );
}
