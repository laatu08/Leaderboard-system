import { useEffect, useState } from "react";
import { FlatList, Text, View, StyleSheet, RefreshControl } from "react-native";
import { theme } from "../../constants/theme";
import { API_BASE_URL } from "@/constants/config";

type Entry = {
  rank: number;
  username: string;
  rating: number;
};

export default function Leaderboard() {
  const RANK_OPTIONS = [3, 5, 10, 50] as const;

  const [rankLimit, setRankLimit] = useState<number>(50);

  const [data, setData] = useState<Entry[]>([]);
  const [refreshing, setRefreshing] = useState(false);

  const fetchData = async () => {
    const res = await fetch(
      `${API_BASE_URL}/leaderboard?rank=${rankLimit}`
    );
    const json = await res.json();
    setData(json);
  };

  useEffect(() => {
    fetchData();

    const interval = setInterval(() => {
      fetchData();
    }, 2000);

    return () => clearInterval(interval);
  }, [rankLimit]);

  const onRefresh = async () => {
    setRefreshing(true);
    await fetchData();
    setRefreshing(false);
  };

  const getRankColor = (rank: number) => {
    if (rank === 1) return theme.gold;
    if (rank === 2) return theme.silver;
    if (rank === 3) return theme.bronze;
    return theme.accent;
  };

  return (
    <View style={{ flex: 1, backgroundColor: theme.bg }}>
      <View style={styles.rankSelector}>
        {RANK_OPTIONS.map((rank) => {
          const active = rank === rankLimit;

          return (
            <Text
              key={rank}
              onPress={() => setRankLimit(rank)}
              style={[styles.rankOption, active && styles.rankOptionActive]}
            >
              Top {rank}
            </Text>
          );
        })}
      </View>

      <FlatList
        contentContainerStyle={styles.list}
        data={data}
        refreshControl={
          <RefreshControl refreshing={refreshing} onRefresh={onRefresh} />
        }
        keyExtractor={(item) => item.username}
        renderItem={({ item }) => (
          <View style={styles.card}>
            <Text style={[styles.rank, { color: getRankColor(item.rank) }]}>
              #{item.rank}
            </Text>

            <View style={styles.info}>
              <Text style={styles.username}>{item.username}</Text>
              <Text style={styles.rating}>Rating: {item.rating}</Text>
            </View>
          </View>
        )}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  list: {
    padding: 12,
    backgroundColor: theme.bg,
  },
  card: {
    flexDirection: "row",
    alignItems: "center",
    backgroundColor: theme.card,
    borderRadius: 14,
    padding: 14,
    marginBottom: 10,
    borderWidth: 1,
    borderColor: theme.border,
  },
  rank: {
    fontSize: 22,
    fontWeight: "800",
    width: 60,
  },
  info: {
    flex: 1,
  },
  username: {
    fontSize: 16,
    fontWeight: "700",
    color: theme.textPrimary,
  },
  rating: {
    marginTop: 4,
    fontSize: 13,
    color: theme.textSecondary,
  },
  rankSelector: {
    flexDirection: "row",
    justifyContent: "space-around",
    paddingHorizontal: 12,
    paddingVertical: 10,
    backgroundColor: theme.bg,
  },

  rankOption: {
    paddingVertical: 6,
    paddingHorizontal: 14,
    borderRadius: 999,
    borderWidth: 1,
    borderColor: theme.border,
    color: theme.textSecondary,
    fontSize: 14,
    fontWeight: "600",
    backgroundColor: theme.card,
  },

  rankOptionActive: {
    backgroundColor: theme.accent,
    color: "#ffffff",
    borderColor: theme.accent,
  },
});
