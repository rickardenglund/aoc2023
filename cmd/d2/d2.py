def main(bag, input_str):
    lines = input_str.splitlines()
    cnt = 0
    game_id = 0
    for game_str in lines:
        game_id += 1
        game_str = game_str[game_str.find(":") + 2:]
        if game_is_possible(bag, game_str):
            cnt += game_id
    print("sum:", cnt)


def game_is_possible(bag, game_str):
    for draw_str in game_str.split("; "):
        draw = parse_draw(draw_str)
        if not possible_draw(bag, draw):
            return False
    return True


def parse_draw(draw_str):
    draw = {}
    for color in draw_str.split(", "):
        parts = color.split(" ")
        draw[parts[1]] = int(parts[0])
    return draw


def possible_draw(bag, draw):
    for k in draw:
        if bag[k] < draw[k]:
            return False
    return True


if __name__ == "__main__":
    real_input = True
    bag = {
        "red": 12,
        "green": 13,
        "blue": 14,
    }
    sample_input = """Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"""
    if real_input:
        f = open("../../inputs/day_2.txt")
        main(bag, f.read())
    else:
        main(bag, sample_input)
