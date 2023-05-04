use program_version::get_program_version;

fn main() {
    let git_describe = match option_env!("GIT_DESCRIBE") {
        Some(git_describe) => git_describe,
        None => {
            println!("cargo:rustc-env=GIT_DESCRIBE=unknown");
            "unknown"
        }
    };

    if option_env!("PROGRAM_VERSION").is_none() {
        println!(
            "cargo:rustc-env=PROGRAM_VERSION={}",
            get_program_version(git_describe)
        );
    }
}
