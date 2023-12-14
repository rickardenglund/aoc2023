def main(input_str):
    lines = input_str.splitlines()
    cnt = 0
    game_id = 0
    for game_str in lines:
        game_id += 1
        game_str = game_str[game_str.find(":") + 2:]
        bag = smallest_possible_bag(game_str)
        cnt += power(bag)

    print("sum:", cnt)


def power(bag):
    pwr = 1
    for v in bag.values():
        pwr *= v
    return pwr


def smallest_possible_bag(game_str):
    min_bag = {}
    for draw_str in game_str.split("; "):
        draw = parse_draw(draw_str)
        for clr in draw:
            min_bag[clr] = max(min_bag.get(clr, 0), draw[clr])
    return min_bag


def parse_draw(draw_str):
    draw = {}
    for color in draw_str.split(", "):
        parts = color.split(" ")
        draw[parts[1]] = int(parts[0])
    return draw


if __name__ == "__main__":
    real_input = True
    sample_input = """Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"""
    if real_input:
        f = open("../../inputs/day_2.txt")
        main(f.read())
    else:
        main(sample_input)
