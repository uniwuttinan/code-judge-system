
import { Button, TablePagination } from "@mui/material";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import { useState } from "react";
import { usePaginationChallenge } from "../swrs/challenge";
import { Challenge } from "../types/challenge";
import { ShowStatusIcon } from "./StatusIcon";


export function ChallengeTable() {
  const [page, setPage] = useState(0);
  const [limit, setLimit] = useState(10);
  const [order, _] = useState("asc");

  const { data, isLoading, isError } = usePaginationChallenge(
    page + 1,
    limit,
    order,
    "id"
  );

  if (isLoading) return <div>Loading...</div>;
  if (isError) return <div>Error</div>;
  if (data.items === null) return null;

  return (
    <Paper sx={{ width: "100%", overflow: "hidden" }}>
      <TableContainer>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Challenge Name</TableCell>
              <TableCell align="right">Created By</TableCell>
              <TableCell align="right">Status</TableCell>
              <TableCell align="right">Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data ? (
              data.items.map((c) => ChallengeTableRow(c))
            ) : (
              <h1>Not Found</h1>
            )}
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        align="right"
        component="div"
        count={data.total}
        page={page}
        rowsPerPage={limit}
        onPageChange={(_, newPage) => {
          setPage(newPage);
        }}
        onRowsPerPageChange={(e) => setLimit(parseInt(e.target.value, 10))}
      />
    </Paper>
  );
}

function ChallengeTableRow(challenge: Challenge) {
  return (
    <TableRow
      key={challenge.challenge_id}
      sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
    >
      <TableCell component="th" scope="row">
        {challenge.name}
      </TableCell>
      <TableCell component="th" scope="row" align="right">
        {challenge.user.DisplayName}
      </TableCell>
      <TableCell align="right">
        {ShowStatusIcon(challenge.submission_status)}
      </TableCell>
      <TableCell align="right">
        <Button
          variant="contained"
          color="primary"
          href={`/solve/${challenge.challenge_id}`}
        >
          Solve
        </Button>
      </TableCell>
    </TableRow>
  );
}
