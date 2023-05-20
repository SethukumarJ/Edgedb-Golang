type User {
    required property email -> str {
        constraint exclusive;
    };
    required property password -> str;
    required property dob -> str;
    required property verified -> bool;
    required property firstname -> std::optional<std::str>;  // Mark firstname as optional
    required property lastname -> str;
    required property phone -> str;
    required property cv -> str;
    property accesstoken -> str;
    property refreshtoken -> str;
}
