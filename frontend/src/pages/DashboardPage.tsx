import {
  Box,
  Button,
  Container,
  CssBaseline,
  Divider,
  Paper,
  Typography,
} from "@mui/material";
import { ChallengeTable } from "../components/ChallengeTable";
import { Navbar } from "../components/Navbar";
import { useUser } from "../contexts/user.provider";

export default function DashboardPage() {
  const { user } = useUser();

  return (
    <Container sx={{ width: "100%" }} disableGutters>
      <CssBaseline />

      <Navbar />

      <Container>
        <Paper sx={{ padding: 3, mt: 15 }}>
          <Box justifyContent="space-between" display="flex">
            <Typography variant="h4" component="h1" align="left">
              Challenge
            </Typography>

            {user && user.role === "ADMIN" ? (
              <Button
                variant="contained"
                color="primary"
                href={`/challenge/create`}
              >
                Create
              </Button>
            ) : null}
          </Box>

          <Divider sx={{ my: 3 }} />

          <ChallengeTable />
        </Paper>
      </Container>
    </Container>
  );
}
