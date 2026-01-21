import { useState } from "react";
import { FlatList, Text, TextInput, View, StyleSheet } from "react-native";
import { theme } from "../../constants/theme";
import { useRef } from "react";
import { API_BASE_URL } from "@/constants/config";

type Result = {
  rank: number;
  username: string;
  rating: number;
};

export default function Search() {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState<Result[]>([]);

  const debounceRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const search = (text: string) => {
    setQuery(text);

    if (debounceRef.current) {
      clearTimeout(debounceRef.current);
    }

    debounceRef.current = setTimeout(async () => {
      if (!text) {
        setResults([]);
        return;
      }

      const res = await fetch(`${API_BASE_URL}/search?query=${text}`);
      const json = await res.json();
      setResults(json);
    }, 300);
  };

  return (
    <View style={styles.container}>
      <TextInput
        placeholder="Search player..."
        placeholderTextColor={theme.textSecondary}
        value={query}
        onChangeText={search}
        style={styles.input}
      />

      <FlatList
        data={results}
        keyExtractor={(item) => item.username}
        renderItem={({ item }) => (
          <View style={styles.card}>
            <Text style={styles.rank}>#{item.rank}</Text>
            <View>
              <Text style={styles.username}>{item.username}</Text>
              <Text style={styles.rating}>{item.rating}</Text>
            </View>
          </View>
        )}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    padding: 12,
    backgroundColor: theme.bg,
    flex: 1,
  },
  input: {
    backgroundColor: theme.card,
    borderRadius: 14,
    padding: 14,
    fontSize: 16,
    borderWidth: 1,
    borderColor: theme.border,
    marginBottom: 12,
    color: theme.textPrimary,
  },
  card: {
    flexDirection: "row",
    alignItems: "center",
    backgroundColor: theme.card,
    padding: 14,
    borderRadius: 14,
    marginBottom: 10,
    borderWidth: 1,
    borderColor: theme.border,
  },
  rank: {
    fontSize: 18,
    fontWeight: "800",
    color: theme.accent,
    width: 50,
  },
  username: {
    fontSize: 16,
    fontWeight: "700",
    color: theme.textPrimary,
  },
  rating: {
    fontSize: 13,
    color: theme.textSecondary,
  },
});
